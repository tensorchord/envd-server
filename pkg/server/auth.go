// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/errdefs"
)

// @Summary     register the user.
// @Description register the user for the given public key.
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       request body     types.AuthNRequest true "query params"
// @Success     200     {object} types.AuthNResponse
// @Router      /register [post]
func (s Server) register(c *gin.Context) error {
	var req types.AuthNRequest
	if err := c.BindJSON(&req); err != nil {
		return NewError(http.StatusInternalServerError, err, "gin.bind-json")
	}

	token, err := s.UserService.Register(c.Request.Context(), req.LoginName, req.Password)
	if err != nil {
		if errdefs.IsConflict(err) {
			return NewError(http.StatusConflict, err, "user.register")
		}
		return NewError(http.StatusInternalServerError, err, "user.register")
	}
	res := types.AuthNResponse{
		LoginName:     req.LoginName,
		IdentityToken: token,
		Status:        types.AuthSuccess,
	}
	c.JSON(http.StatusOK, res)
	return nil
}

// @Summary     login the user.
// @Description login to the server.
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       request body     types.AuthNRequest true "query params"
// @Success     200     {object} types.AuthNResponse
// @Router      /login [post]
func (s Server) login(c *gin.Context) error {
	var req types.AuthNRequest
	if err := c.BindJSON(&req); err != nil {
		return NewError(http.StatusInternalServerError, err, "gin.bind-json")
	}

	succeeded, token, err := s.UserService.Login(c.Request.Context(),
		req.LoginName, req.Password, s.Auth)
	if err != nil {
		return NewError(http.StatusUnauthorized, err, "user.login")
	}
	if !succeeded {
		return NewError(http.StatusUnauthorized, err, "user.login")
	}

	res := types.AuthNResponse{
		LoginName:     req.LoginName,
		IdentityToken: token,
		Status:        types.AuthSuccess,
	}
	c.JSON(200, res)
	return nil
}
