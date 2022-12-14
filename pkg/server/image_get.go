// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/errdefs"
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
	owner := c.GetString(ContextLoginName)

	var req types.ImageGetRequest
	if err := c.BindUri(&req); err != nil {
		return NewError(http.StatusInternalServerError, err, "gin.bind-json")
	}

	meta, err := s.ImageService.GetImage(c.Request.Context(), owner, req.Name)
	if err != nil {
		if errdefs.IsNotFound(err) {
			return NewError(http.StatusNotFound, err, "image.get")
		} else if errdefs.IsUnauthorized(err) {
			return NewError(http.StatusUnauthorized, err, "image.get")
		}
		return NewError(http.StatusInternalServerError, err, "image.get")
	}
	if meta == nil {
		return NewError(http.StatusNotFound,
			errors.New("meta is nil"), "image.get")
	}

	c.JSON(http.StatusOK, types.ImageGetResponse{
		ImageMeta: *meta,
	})
	return nil
}
