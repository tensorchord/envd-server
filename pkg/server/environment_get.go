// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/consts"
)

// @Summary     Get the environment.
// @Description Get the environment with the given environment name.
// @Tags        environment
// @Accept      json
// @Produce     json
// @Param       login_name     path     string true "login name" example("alice")
// @Param       name           path     string true "environment name" example("pytorch-example")
// @Success     200            {object} types.EnvironmentGetResponse
// @Router      /users/{login_name}/environments/{name} [get]
func (s *Server) environmentGet(c *gin.Context) {
	it := c.GetString(ContextLoginName)

	var req types.EnvironmentGetRequest
	if err := c.BindUri(&req); err != nil {
		c.JSON(500, err)
		return
	}

	pod, err := s.Client.CoreV1().Pods("default").Get(c, req.Name, metav1.GetOptions{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, types.EnvironmentGetResponse{})
			return
		}
		c.JSON(500, err)
		return
	}
	if pod.Labels[consts.PodLabelUID] != it {
		logrus.WithFields(logrus.Fields{
			"loginname_in_pod":     pod.Labels[consts.PodLabelUID],
			"loginname_in_request": it,
		}).Debug("mismatch loginname")
		respondWithError(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	if pod == nil {
		c.JSON(http.StatusNotFound, types.EnvironmentGetResponse{})
		return
	}

	e, err := generateEnvironmentFromPod(*pod)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(http.StatusOK, types.EnvironmentGetResponse{
		Environment: e,
	})
}
