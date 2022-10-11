// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/tensorchord/envd-server/api/types"
)

func (s *Server) environmentList(c *gin.Context) {
	var req types.EnvironmentListRequest
	if err := c.BindUri(&req); err != nil {
		logrus.Error(err, "failed to bind URI")
		c.JSON(500, err)
		return
	}
	pod, err := s.client.CoreV1().Pods(
		"default").Get(c, req.IdentityToken, metav1.GetOptions{})
	if err != nil {
		logrus.WithField("req", req).Error(err)
		c.JSON(500, err)
		return
	}

	res := types.EnvironmentListResponse{
		Pod: *pod,
	}
	c.JSON(200, res)
}
