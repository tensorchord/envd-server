// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tensorchord/envd-server/api/types"
)

// @Summary     Create the environment.
// @Description Create the environment.
// @Security    Authentication
// @Tags        environment
// @Accept      json
// @Produce     json
// @Param       login_name     path     string true "login name" example("alice")
// @Param       request        body     types.EnvironmentCreateRequest true "query params"
// @Success     201            {object} types.EnvironmentCreateResponse
// @Router      /users/{login_name}/environments [post]
func (s Server) environmentCreate(c *gin.Context) error {
	owner := c.GetString(ContextLoginName)

	var req types.EnvironmentCreateRequest
	if err := c.BindJSON(&req); err != nil {
		return NewError(http.StatusInternalServerError, err, "gin.bind-json")
	}

	// Create the image in DB if it does not exist.
	meta, err := s.ImageService.CreateImageIfNotExist(
		c.Request.Context(), owner, req.Spec.Image)
	if err != nil {
		return NewError(http.StatusInternalServerError,
			err, "image-service.create-image")
	}
	if meta == nil {
		return NewError(http.StatusInternalServerError,
			errors.New("meta is nil"), "image-service.create-image")
	}

	// Create the environment.
	env, err := s.Runtime.EnvironmentCreate(c.Request.Context(),
		owner, req.Environment, *meta)
	if err != nil {
		return NewError(http.StatusInternalServerError,
			err, "runtime.create-environment")
	}

	resp := types.EnvironmentCreateResponse{
		Created: *env,
	}
	c.JSON(http.StatusCreated, resp)
	return nil
}
