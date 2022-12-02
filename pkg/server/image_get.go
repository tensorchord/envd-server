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
	"github.com/sirupsen/logrus"

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
func (s *Server) imageGet(c *gin.Context) {
	it := c.GetString(ContextLoginName)

	var req types.ImageGetRequest
	if err := c.BindUri(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	name, err := url.PathUnescape(req.Name)
	if err != nil {
		logrus.Infof("cannot unescape the requested image name: %s", req.Name)
		c.JSON(http.StatusBadRequest, err)
	}

	imageInfo, err := s.Queries.GetImageInfo(context.Background(), query.GetImageInfoParams{OwnerToken: it, Name: name})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusBadRequest, errors.Newf("cannot find the image(%s)", req.Name))
			return
		} else {
			logrus.Warnf("cannot get the image info: %+v", err)
			c.JSON(http.StatusInternalServerError, "internal error")
			return
		}
	}
	meta, err := util.DaoToImageMeta(imageInfo)
	if err != nil {
		logrus.Warnf("cannot get label info: %+v", err)
		c.JSON(http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, types.ImageGetResponse{
		ImageMeta: *meta,
	})

}
