// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/query"
	"github.com/tensorchord/envd-server/pkg/util"
)

// @Summary     Get the image.
// @Description Get the image with the given image name.
// @Security    Authentication
// @Tags        image
// @Accept      json
// @Produce     json
// @Param       login_name     path     string true "login name" example("alice")
// @Param       name           path     string true "image name" example("pytorch-example")
// @Success     200            {object} types.ImageGetResponse
// @Router      /users/{login_name}/images/{name} [get]
func (s *Server) imageGet(c *gin.Context) error {
	it := c.GetString(ContextLoginName)

	var req types.ImageGetRequest
	if err := c.BindUri(&req); err != nil {
		return NewError(http.StatusInternalServerError, err, "gin.bind-json")
	}

	name, err := url.PathUnescape(req.Name)
	if err != nil {
		return NewError(http.StatusBadRequest, err, "url.path-unescape")
	}

	imageInfo, err := s.Queries.GetImageInfo(context.Background(), query.GetImageInfoParams{OwnerToken: it, Name: name})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return NewError(http.StatusNotFound, err, "db.get-image")
		} else {
			return NewError(http.StatusInternalServerError, err, "db.get-image")
		}
	}
	meta, err := util.DaoToImageMeta(imageInfo)
	if err != nil {
		return NewError(http.StatusInternalServerError, err, "db.get-image")
	}
	c.JSON(http.StatusOK, types.ImageGetResponse{
		ImageMeta: *meta,
	})
	return nil
}
