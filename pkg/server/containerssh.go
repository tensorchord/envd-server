// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"crypto/subtle"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.containerssh.io/libcontainerssh/auth"
	"go.containerssh.io/libcontainerssh/config"
	"golang.org/x/crypto/ssh"

	"github.com/tensorchord/envd-server/sshname"
)

// @Summary     Update the config of containerssh.
// @Description It is called by the containerssh webhook. and is not expected to be used externally.
// @Tags        ssh-internal
// @Accept      json
// @Produce     json
// @Param       request body config.Request true "query params"
// @Success     200
// @Router      /config [post]
func (s Server) OnConfig(c *gin.Context) error {
	var req config.Request
	if err := c.BindJSON(&req); err != nil {
		return NewError(http.StatusInternalServerError, err, "gin.bind-json")
	}

	_, name, err := sshname.GetInfo(req.Username)
	if err != nil {
		return NewError(http.StatusInternalServerError, err, "sshname.get-info")
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
	return nil
}

// @Summary     authenticate the public key.
// @Description It is called by the containerssh webhook. and is not expected to be used externally.
// @Tags        ssh-internal
// @Accept      json
// @Produce     json
// @Param       request body     auth.PublicKeyAuthRequest true "query params"
// @Success     200     {object} auth.ResponseBody
// @Router      /pubkey [post]
func (s Server) OnPubKey(c *gin.Context) error {
	var req auth.PublicKeyAuthRequest
	if err := c.BindJSON(&req); err != nil {
		return NewError(http.StatusInternalServerError, err, "gin.bind-json")
	}

	owner, _, err := sshname.GetInfo(req.Username)
	if err != nil {
		return NewError(http.StatusInternalServerError, err, "sshname.get-info")
	}

	skey, err := s.UserService.GetPubKey(owner)
	if err != nil {
		return NewError(http.StatusInternalServerError, err, "db.get-pubkey-from-user")
	}
	if skey == nil {
		return NewError(http.StatusInternalServerError, errors.New("key is not found"), "db.get-pubkey-from-user")
	}
	key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(req.PublicKey.PublicKey))
	if err != nil {
		return NewError(http.StatusInternalServerError, err, "ssh.parse-auth-key")
	}
	if subtle.ConstantTimeCompare(key.Marshal(), skey) == 1 {
		logrus.WithFields(logrus.Fields{
			"username":    req.Username,
			"remote-addr": req.RemoteAddress,
		}).Debug("auth success")
		res := auth.ResponseBody{
			Success: true,
		}
		c.JSON(http.StatusOK, res)
		return nil
	}

	logrus.WithFields(logrus.Fields{
		"username":    req.Username,
		"remote-addr": req.RemoteAddress,
	}).Debug("auth failed")
	res := auth.ResponseBody{
		Success: false,
	}
	c.JSON(http.StatusOK, res)
	return nil
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
