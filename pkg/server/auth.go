// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/query"
)

type AuthInfo struct {
	IdentityToken string
	PublicKey     ssh.PublicKey
}

// @Summary     authenticate the user.
// @Description authenticate the user for the given public key.
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       request body     types.AuthRequest true "query params"
// @Success     200     {object} types.AuthResponse
// @Router      /auth [post]
func (s *Server) auth(c *gin.Context) {
	var req types.AuthRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(500, err)
		return
	}

	key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(req.PublicKey))
	if err != nil {
		c.JSON(500, err)
		return
	}
	_, err = s.Queries.CreateUser(context.Background(), query.CreateUserParams{IdentityToken: req.IdentityToken, PublicKey: key.Marshal()})
	if err != nil {
		c.JSON(500, err)
		return
	}
	res := types.AuthResponse{
		IdentityToken: req.IdentityToken,
		Status:        "login succeeded",
	}
	c.JSON(200, res)
}
