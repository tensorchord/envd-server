// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"context"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/util"
)

// @Summary     List the images.
// @Description List the images.
// @Security    Authentication
// @Tags        image
// @Accept      json
// @Produce     json
// @Param       login_name     path     string true "login name" example("alice")
// @Param       name           path     string true "image name" example("pytorch-example")
// @Success     200            {object} types.ImageListResponse
// @Router      /users/{login_name}/images [get]
func (s *Server) imageList(c *gin.Context) error {
	it := c.GetString(ContextLoginName)

	resp := types.ImageListResponse{}
	images, err := s.Queries.ListImageByOwner(context.Background(), it)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// No image found
			c.JSON(http.StatusOK, resp)
			return nil
		} else {
			return NewError(http.StatusInternalServerError, err, "db.list-image")
		}
	}
	for _, info := range images {
		item, err := util.DaoToImageMeta(info)
		if err != nil {
			return NewError(http.StatusInternalServerError, err, "db.list-image")
		}
		resp.Items = append(resp.Items, *item)
	}
	c.JSON(http.StatusOK, resp)
	return nil
}
