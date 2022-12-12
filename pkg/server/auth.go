// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/service/user"
)

// @Summary     register the user.
// @Description register the user for the given public key.
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       request body     types.AuthNRequest true "query params"
// @Success     200     {object} types.AuthNResponse
// @Router      /register [post]
func (s *Server) register(c *gin.Context) {
	var req types.AuthNRequest
	if err := c.BindJSON(&req); err != nil {
		logrus.Debug("failed to bind json", err)
		c.JSON(500, err)
		return
	}

	key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(req.PublicKey))
	if err != nil {
		logrus.Debug("failed to parse auth key", err)
		c.JSON(500, err)
		return
	}
	userService := user.NewService(s.Queries,
		s.JWTSecret, s.JWTExpirationTimeout)
	token, err := userService.Register(req.LoginName, req.Password, key.Marshal())
	if err != nil {
		logrus.Warnf("register error: %+v", err)
		c.JSON(500, err)
		return
	}
	res := types.AuthNResponse{
		LoginName:     req.LoginName,
		IdentityToken: token,
		Status:        types.AuthSuccess,
	}
	c.JSON(200, res)
}

// @Summary     login the user.
// @Description login to the server.
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       request body     types.AuthNRequest true "query params"
// @Success     200     {object} types.AuthNResponse
// @Router      /login [post]
func (s *Server) login(c *gin.Context) {
	var req types.AuthNRequest
	if err := c.BindJSON(&req); err != nil {
		logrus.Debug("failed to bind json", err)
		c.JSON(500, err)
		return
	}

	userService := user.NewService(s.Queries,
		s.JWTSecret, s.JWTExpirationTimeout)
	succeeded, token, err := userService.Login(req.LoginName, req.Password, s.Auth)
	if err != nil {
		logrus.Debug("login error: ", err)
		respondWithError(c, http.StatusUnauthorized,
			fmt.Sprintf("auth failed: %+v", err))
		return
	}
	if !succeeded {
		respondWithError(c, http.StatusUnauthorized,
			fmt.Sprintf("auth failed: %+v", err))
		return
	}
	res := types.AuthNResponse{
		LoginName:     req.LoginName,
		IdentityToken: token,
		Status:        types.AuthSuccess,
	}
	c.JSON(200, res)
}
