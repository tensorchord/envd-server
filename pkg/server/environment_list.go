// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/consts"
	"github.com/tensorchord/envd-server/pkg/util/imageutil"
)

// @Summary List the environment.
// @Description List the environment.
// @Tags environment
// @Accept json
// @Produce json
// @Param identity_token path string true "identity token" example("a332139d39b89a241400013700e665a3")
// @Success 200 {object} types.EnvironmentListResponse
// @Router /users/{identity_token}/environments [get]
func (s *Server) environmentList(c *gin.Context) {
	it := c.GetString("identity_token")

	ls := labels.Set{
		consts.LabelUID: it,
	}

	pods, err := s.client.CoreV1().Pods(
		"default").List(c, metav1.ListOptions{
		LabelSelector: ls.String(),
	})
	if err != nil {
		logrus.WithField("identity_token", it).Error(err)
		if k8serrors.IsNotFound(err) {
			c.JSON(404, types.EnvironmentListResponse{})
			return
		}
		c.JSON(500, err)
		return
	}

	res := types.EnvironmentListResponse{
		Items: []types.Environment{},
	}

	for _, p := range pods.Items {
		e, err := generateEnvironmentFromPod(p)
		if err != nil {
			logrus.WithField("identity_token", it).Error(err)
			c.JSON(500, err)
			return
		}
		res.Items = append(res.Items, e)
	}
	c.JSON(200, res)
}

// nolint:unparam
func generateEnvironmentFromPod(p v1.Pod) (types.Environment, error) {
	e := types.Environment{
		// TODO(gaocegege): Filter non `envd.tensorchord.ai/` labels
		ObjectMeta: types.ObjectMeta{
			Labels: p.Annotations,
			Name:   p.Name,
		},
		Spec: types.EnvironmentSpec{},
	}
	if len(p.Spec.Containers) != 0 {
		e.Spec.Image = p.Spec.Containers[0].Image
	}

	ports, err := imageutil.PortsFromLabel(p.Annotations[consts.ImageLabelPorts])
	if err != nil {
		return e, err
	}
	e.Spec.Ports = ports

	e.Status.Phase = string(p.Status.Phase)
	return e, nil
}
