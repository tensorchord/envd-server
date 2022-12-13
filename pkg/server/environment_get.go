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
func (s Server) environmentGet(c *gin.Context) error {
	owner := c.GetString(ContextLoginName)

	var req types.EnvironmentGetRequest
	if err := c.BindUri(&req); err != nil {
		return NewError(http.StatusInternalServerError, err, "gin.bind-json")
	}

	e, err := s.Runtime.EnvironmentGet(c.Request.Context(), owner, req.Name)
	if err != nil {
		if errdefs.IsNotFound(err) {
			return NewError(http.StatusNotFound, err, "runtime.get-environment")
		} else if errdefs.IsUnauthorized(err) {
			return NewError(http.StatusUnauthorized, err, "runtime.get-environment")
		}
		return NewError(http.StatusInternalServerError, err, "runtime.get-environment")
	}

	c.JSON(http.StatusOK, types.EnvironmentGetResponse{
		Environment: *e,
	})
	return nil
}
