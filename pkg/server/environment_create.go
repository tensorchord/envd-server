// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgtype"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/image"
	"github.com/tensorchord/envd-server/pkg/query"
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
	it := c.GetString(ContextLoginName)

	var req types.EnvironmentCreateRequest
	if err := c.BindJSON(&req); err != nil {
		return NewError(http.StatusInternalServerError, err, "gin.bind-json")
	}

	meta, err := image.FetchMetadata(c.Request.Context(), req.Spec.Image)
	if err != nil {
		return NewError(http.StatusInternalServerError, err, "image.fetch-metadata")
	}
	var pglabel pgtype.JSONB
	err = pglabel.Set(meta.Labels)
	if err != nil {
		return NewError(http.StatusInternalServerError, err, "pglabel.set")
	}
	_, err = s.Queries.CreateImageInfo(context.Background(),
		query.CreateImageInfoParams{OwnerToken: it,
			Name: meta.Name, Digest: meta.Digest, Created: meta.Created, Size: meta.Size, Labels: pglabel})
	if err != nil {
		return NewError(http.StatusInternalServerError, err, "db.create-image")
	}

	env, err := s.Runtime.EnvironmentCreate(c.Request.Context(), it, req.Environment, meta)
	if err != nil {
		return NewError(http.StatusInternalServerError, err, "runtime.create-environment")
	}

	resp := types.EnvironmentCreateResponse{
		Created: *env,
	}
	c.JSON(http.StatusCreated, resp)
	return nil
}
