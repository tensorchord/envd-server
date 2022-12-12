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

// @Summary     Get the environment.
// @Description Get the environment with the given environment name.
// @Security    Authentication
// @Tags        environment
// @Accept      json
// @Produce     json
// @Param       login_name     path     string true "login name" example("alice")
// @Param       name           path     string true "environment name" example("pytorch-example")
// @Success     200            {object} types.EnvironmentGetResponse
// @Router      /users/{login_name}/environments/{name} [get]
func (s *Server) environmentGet(c *gin.Context) {
	owner := c.GetString(ContextLoginName)

	var req types.EnvironmentGetRequest
	if err := c.BindUri(&req); err != nil {
		c.JSON(500, err)
		return
	}

	e, err := s.Runtime.EnvironmentGet(c.Request.Context(), owner, req.Name)
	if err != nil {
		if errdefs.IsNotFound(err) {
			c.JSON(http.StatusNotFound, err)
			return
		} else if errdefs.IsUnauthorized(err) {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, types.EnvironmentGetResponse{
		Environment: *e,
	})
}
