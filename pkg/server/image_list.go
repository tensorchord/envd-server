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
	"github.com/sirupsen/logrus"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/util"
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
	images, err := s.Queries.ListImageByOwner(context.Background(), it)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// No image found
			c.JSON(http.StatusOK, resp)
			return
		} else {
			logrus.Warnf("cannot get the image info: %+v", err)
			c.JSON(http.StatusInternalServerError, "internal error")
			return
		}
	}
	for _, info := range images {
		item, err := util.DaoToImageMeta(info)
		if err != nil {
			logrus.Warnf("cannot convert dao to image info: %+v", err)
			c.JSON(http.StatusInternalServerError, "internal error")
			return
		}
		resp.Items = append(resp.Items, *item)
	}
	c.JSON(http.StatusOK, resp)
}
