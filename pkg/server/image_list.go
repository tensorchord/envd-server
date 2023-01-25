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

// @Summary     List the images.
// @Description List the images.
// @Security    Authentication
// @Tags        image
// @Accept      json
// @Produce     json
// @Param       login_name path     string true "login name" example("alice")
// @Success     200        {object} types.ImageListResponse
// @Router      /users/{login_name}/images [get]
func (s *Server) imageList(c *gin.Context) error {
	owner := c.GetString(ContextLoginName)

	resp := types.ImageListResponse{}
	images, err := s.ImageService.ListImages(c.Request.Context(), owner)
	if err != nil {
		if errdefs.IsNotFound(err) {
			return NewError(http.StatusNotFound, err, "image.list")
		} else {
			return NewError(http.StatusInternalServerError, err, "image.list")
		}
	}

	resp.Items = images
	c.JSON(http.StatusOK, resp)
	return nil
}
