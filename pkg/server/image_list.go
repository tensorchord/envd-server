// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tensorchord/envd-server/api/types"
)

// @Summary     List the images.
// @Description List the images.
// @Tags        image
// @Accept      json
// @Produce     json
// @Param       identity_token path     string true "identity token" example("a332139d39b89a241400013700e665a3")
// @Param       name           path     string true "image name" example("pytorch-example")
// @Success     200            {object} types.ImageListResponse
// @Router      /users/{identity_token}/images [get]
func (s *Server) imageList(c *gin.Context) {
	it := c.GetString("identity_token")

	resp := types.ImageListResponse{}
	for _, info := range s.imageInfo {
		if info.OwnerToken == it {
			resp.Items = append(resp.Items, info.ImageMeta)
		}
	}
	c.JSON(http.StatusOK, resp)
}
