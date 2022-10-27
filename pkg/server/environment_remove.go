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
// @Tags        environment
// @Accept      json
// @Produce     json
// @Param       identity_token path     string true "identity token" example("a332139d39b89a241400013700e665a3")
// @Param       name           path     string true "environment name" example("pytorch-example")
// @Success     200            {object} types.EnvironmentRemoveResponse
// @Router      /users/{identity_token}/environments/{name} [delete]
func (s *Server) environmentRemove(c *gin.Context) {
	it := c.GetString("identity_token")

	var req types.EnvironmentRemoveRequest
	if err := c.BindUri(&req); err != nil {
		c.JSON(500, err)
		return
	}

	logger := logrus.WithFields(logrus.Fields{
		"name":           req.Name,
		"identity_token": it,
	})
	pod, err := s.client.CoreV1().Pods("default").Get(c, req.Name, metav1.GetOptions{})
	if !k8serrors.IsNotFound(err) {
		if err != nil {
			logger.Error(err)
			c.JSON(500, err)
			return
		}
		if pod.Labels[consts.LabelUID] != it {
			logger.WithFields(logrus.Fields{
				"identity_token_in_pod":     pod.Labels[consts.LabelUID],
				"identity_token_in_request": it,
			}).Debug("mismatch identity_token")
			respondWithError(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		err = s.client.CoreV1().Pods(
			"default").Delete(c, req.Name, metav1.DeleteOptions{})
		if err != nil && !k8serrors.IsNotFound(err) {
			logger.Error(err)
			c.JSON(500, err)
			return
		}
		logger.Debugf("pod %s is deleted", req.Name)
	}

	service, err := s.client.CoreV1().Services("default").Get(c, req.Name, metav1.GetOptions{})
	if !k8serrors.IsNotFound(err) {
		if err != nil {
			logger.Error(err)
			c.JSON(500, err)
			return
		}
		if service.Labels[consts.LabelUID] != it {
			logger.WithFields(logrus.Fields{
				"identity_token_in_pod":     pod.Labels[consts.LabelUID],
				"identity_token_in_request": it,
			}).Debug("mismatch identity_token")
			respondWithError(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		err = s.client.CoreV1().Services("default").Delete(c, req.Name, metav1.DeleteOptions{})
		if err != nil && !k8serrors.IsNotFound(err) {
			logger.Error(err)
			c.JSON(500, err)
			return
		}
		logger.Debugf("service %s is deleted", req.Name)
	}

	c.JSON(200, types.EnvironmentRemoveResponse{})
}
