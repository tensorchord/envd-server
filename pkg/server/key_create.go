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

// @Summary     Create the key.
// @Description Create the key.
// @Security    Authentication
// @Tags        key
// @Accept      json
// @Produce     json
// @Param       login_name     path     string true "login name" example("alice")
// @Param       request        body     types.KeyCreateRequest true "query params"
// @Success     201            {object} types.KeyCreateResponse
// @Router      /users/{login_name}/keys [post]
func (s Server) keyCreate(c *gin.Context) error {
	owner := c.GetString(ContextLoginName)

	var req types.KeyCreateRequest
	if err := c.BindJSON(&req); err != nil {
		return NewError(http.StatusInternalServerError, err, "gin.bind-json")
	}

	if err := s.UserService.CreatePubKey(c.Request.Context(),
		owner, req.Name, []byte(req.PublicKey)); err != nil {
		if errdefs.IsConflict(err) {
			return NewError(http.StatusConflict, err, "user.create-pubkey")
		}
		return NewError(http.StatusInternalServerError, err, "user.create-pubkey")
	}

	c.JSON(http.StatusOK, types.KeyCreateResponse{
		Name:      req.Name,
		PublicKey: req.PublicKey,
		LoginName: owner,
	})
	return nil
}
