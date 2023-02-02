// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/errdefs"
)

// @Summary     List the environment.
// @Description List the environment.
// @Security    Authentication
// @Tags        environment
// @Accept      json
// @Produce     json
// @Param       login_name path     string true "login name" example("alice")
// @Success     200        {object} types.EnvironmentListResponse
// @Router      /users/{login_name}/environments [get]
func (s Server) environmentList(c *gin.Context) error {
	owner := c.GetString(ContextLoginName)
	logger := logrus.WithField(ContextLoginName, owner)

	items, err := s.Runtime.EnvironmentList(c.Request.Context(), owner)
	if err != nil {
		if errdefs.IsNotFound(err) {
			return NewError(http.StatusNotFound, err, "runtime.list-environment")
		}
		return NewError(http.StatusInternalServerError, err, "runtime.list-environment")
	}

	res := types.EnvironmentListResponse{
		Items: items,
	}

	logger.WithField("count", len(res.Items)).
		Debug("list the environments successfully")
	c.JSON(http.StatusOK, res)
	return nil
}
