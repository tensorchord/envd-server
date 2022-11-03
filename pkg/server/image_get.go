// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"net/http"
	"net/url"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/tensorchord/envd-server/api/types"
)

// @Summary     Get the image.
// @Description Get the image with the given image name.
// @Tags        image
// @Accept      json
// @Produce     json
// @Param       identity_token path     string true "identity token" example("a332139d39b89a241400013700e665a3")
// @Param       name           path     string true "image name" example("pytorch-example")
// @Success     200            {object} types.ImageGetResponse
// @Router      /users/{identity_token}/images/{name} [get]
func (s *Server) imageGet(c *gin.Context) {
	it := c.GetString("identity_token")

	var req types.ImageGetRequest
	if err := c.BindUri(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	name, err := url.PathUnescape(req.Name)
	if err != nil {
		logrus.Info("cannot unescape the requested image name: %s", req.Name)
		c.JSON(http.StatusBadRequest, err)
	}
	for _, info := range s.imageInfo {
		if info.OwnerToken == it && info.Name == name {
			c.JSON(http.StatusOK, types.ImageGetResponse{
				ImageMeta: info.ImageMeta,
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, errors.Newf("cannot find the image(%s)", req.Name))
}
