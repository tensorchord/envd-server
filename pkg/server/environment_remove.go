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

// @Summary     Remove the environment.
// @Description Remove the environment.
// @Security    Authentication
// @Tags        environment
// @Accept      json
// @Produce     json
// @Param       login_name     path     string true "login name" example("alice")
// @Param       name           path     string true "environment name" example("pytorch-example")
// @Success     200            {object} types.EnvironmentRemoveResponse
// @Router      /users/{login_name}/environments/{name} [delete]
func (s *Server) environmentRemove(c *gin.Context) {
	it := c.GetString(ContextLoginName)

	var req types.EnvironmentRemoveRequest
	if err := c.BindUri(&req); err != nil {
		c.JSON(500, err)
		return
	}

	logger := logrus.WithFields(logrus.Fields{
		"name":           req.Name,
		ContextLoginName: it,
	})
	pod, err := s.Client.CoreV1().Pods("default").Get(c, req.Name, metav1.GetOptions{})
	if !k8serrors.IsNotFound(err) {
		if err != nil {
			logger.Error(err)
			c.JSON(500, err)
			return
		}
		if pod.Labels[consts.PodLabelUID] != it {
			logger.WithFields(logrus.Fields{
				"loginname_in_pod":     pod.Labels[consts.PodLabelUID],
				"loginname_in_request": it,
			}).Debug("mismatch loginname")
			respondWithError(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		err = s.Client.CoreV1().Pods(
			"default").Delete(c, req.Name, metav1.DeleteOptions{})
		if err != nil && !k8serrors.IsNotFound(err) {
			logger.Error(err)
			c.JSON(500, err)
			return
		}
		logger.Debugf("pod %s is deleted", req.Name)
	}

	service, err := s.Client.CoreV1().Services("default").Get(c, req.Name, metav1.GetOptions{})
	if !k8serrors.IsNotFound(err) {
		if err != nil {
			logger.Error(err)
			c.JSON(500, err)
			return
		}
		if service.Labels[consts.PodLabelUID] != it {
			logger.WithFields(logrus.Fields{
				"loginname_in_pod":     pod.Labels[consts.PodLabelUID],
				"loginname_in_request": it,
			}).Debug("mismatch loginname")
			respondWithError(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		err = s.Client.CoreV1().Services("default").Delete(c, req.Name, metav1.DeleteOptions{})
		if err != nil && !k8serrors.IsNotFound(err) {
			logger.Error(err)
			c.JSON(500, err)
			return
		}
		logger.Debugf("service %s is deleted", req.Name)
	}

	c.JSON(200, types.EnvironmentRemoveResponse{})
}
