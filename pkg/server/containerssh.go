// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"crypto/subtle"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.containerssh.io/libcontainerssh/auth"
	"go.containerssh.io/libcontainerssh/config"
	"golang.org/x/crypto/ssh"

	"github.com/tensorchord/envd-server/sshname"
)

// @Summary     Update the config of containerssh.
// @Description It is called by the containerssh webhook. It is not expected to be used externally.
// @Tags        ssh-internal
// @Accept      json
// @Produce     json
// @Param       request body config.Request true "query params"
// @Success     200
// @Router      /config [post]
func (s Server) OnConfig(c *gin.Context) {
	var req config.Request
	if err := c.BindJSON(&req); err != nil {
		logrus.Debugf("gin.bind err: %v", err)
		c.Status(http.StatusBadRequest)
		return
	}

	_, name, err := sshname.GetInfo(req.Username)
	if err != nil {
		logrus.Debugf("sshname.get err: %v", err)
		c.Status(http.StatusBadRequest)
		return
	}

	cfg := config.AppConfig{
		Backend: "sshproxy",
		SSHProxy: config.SSHProxyConfig{
			Server:   name,
			Port:     2222,
			Username: "envd",
		},
	}
	fingerprints := s.serverFingerPrints
	cfg.SSHProxy.AllowedHostKeyFingerprints = fingerprints
	res := config.ResponseBody{
		Config: cfg,
	}
	c.JSON(http.StatusOK, res)
}

// @Summary     authenticate the public key.
// @Description It is called by the containerssh webhook. It is not expected to be used externally.
// @Tags        ssh-internal
// @Accept      json
// @Produce     json
// @Param       request body     auth.PublicKeyAuthRequest true "query params"
// @Success     200     {object} auth.ResponseBody
// @Router      /pubkey [post]
func (s Server) OnPubKey(c *gin.Context) {
	var req auth.PublicKeyAuthRequest
	if err := c.BindJSON(&req); err != nil {
		logrus.Debugf("bind.json err: %v", err)
		c.JSON(http.StatusBadRequest, auth.ResponseBody{Success: false})
		return
	}

	owner, _, err := sshname.GetInfo(req.Username)
	if err != nil {
		logrus.Debugf("sshname.get-info err: %v", err)
		c.JSON(http.StatusBadRequest, auth.ResponseBody{Success: false})
		return
	}

	skeys, err := s.UserService.ListPubKeys(c.Request.Context(), owner)
	if err != nil {
		logrus.Debugf("db.get-pubkey err: %v", err)
		c.JSON(http.StatusBadRequest, auth.ResponseBody{Success: false})
		return
	}
	if len(skeys) == 0 {
		logrus.Debugf("db.get-pubkey err: %v", err)
		c.JSON(http.StatusBadRequest, auth.ResponseBody{Success: false})
		return
	}
	key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(req.PublicKey.PublicKey))
	if err != nil {
		logrus.Debugf("ssh.parse err: %v", err)
		c.JSON(http.StatusBadRequest, auth.ResponseBody{Success: false})
		return
	}

	for _, skey := range skeys {
		logger := logrus.WithFields(logrus.Fields{
			"username":    req.Username,
			"remote-addr": req.RemoteAddress,
			"key-name":    skey.Name,
		})
		if subtle.ConstantTimeCompare(key.Marshal(), skey.PublicKey) == 1 {
			logger.Debug("auth success")
			res := auth.ResponseBody{
				Success: true,
			}
			c.JSON(http.StatusOK, res)
			return
		} else {
			logger.Debug("trying next ssh key")
		}
	}

	logrus.WithFields(logrus.Fields{
		"username":    req.Username,
		"remote-addr": req.RemoteAddress,
	}).Debug("auth failed")
	res := auth.ResponseBody{
		Success: false,
	}
	c.JSON(http.StatusOK, res)
}
