// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/tensorchord/envd-server/api/types"
)

// @Summary     Remove the environment.
// @Description Remove the environment.
// @Security    Authentication
// @Tags        environment
// @Accept      json
// @Produce     json
// @Param       login_name     path     string true "login name" example("alice")
// @Param       name           path     string true "environment name" example("pytorch-example")
// @Success     200            {object} types.EnvironmentRemoveResponse
// @Router      /users/{login_name}/environments/{name} [delete]
func (s *Server) environmentRemove(c *gin.Context) {
	owner := c.GetString(ContextLoginName)

	var req types.EnvironmentRemoveRequest
	if err := c.BindUri(&req); err != nil {
		c.JSON(500, err)
		return
	}

	if err := s.Runtime.EnvironmentRemove(c.Request.Context(), owner, req.Name); err != nil {
		if k8serrors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, err)
			return
		} else if k8serrors.IsUnauthorized(err) {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, types.EnvironmentRemoveResponse{})
}
