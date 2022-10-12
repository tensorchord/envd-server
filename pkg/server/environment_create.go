// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/tensorchord/envd-server/api/types"
)

// @Summary Create the environment.
// @Description Create the environment.
// @Tags environment
// @Accept json
// @Produce json
// @Param request body types.EnvironmentCreateRequest true "query params"
// @Success 200 {object} types.EnvironmentCreateResponse
// @Router /environments [post]
func (s *Server) environmentCreate(c *gin.Context) {
	var req types.EnvironmentCreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(500, err)
		return
	}
	hostKeyPath := "/var/envd/hostkey"
	authKeyPath := "/var/envd/authkey"
	var defaultPermMode int32 = 0600
	expectedPod := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.IdentityToken,
			Namespace: "default",
			Labels: map[string]string{
				"name": req.IdentityToken,
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "envd",
					Image: req.Image,
					Ports: []v1.ContainerPort{
						{
							Name:          "ssh",
							ContainerPort: 2222,
						},
					},
					Env: []v1.EnvVar{
						{
							Name:  "ENVD_HOST_KEY",
							Value: hostKeyPath,
						},
						{
							Name:  "ENVD_AUTHORIZED_KEYS_PATH",
							Value: authKeyPath,
						},
					},
					VolumeMounts: []v1.VolumeMount{
						{
							Name:      "secret",
							ReadOnly:  true,
							MountPath: hostKeyPath,
							SubPath:   "hostkey",
						},
						{
							Name:      "secret",
							ReadOnly:  true,
							MountPath: authKeyPath,
							SubPath:   "publickey",
						},
					},
				},
			},
			Volumes: []v1.Volume{
				{
					Name: "secret",
					VolumeSource: v1.VolumeSource{
						Secret: &v1.SecretVolumeSource{
							SecretName:  "containerssh-secret",
							DefaultMode: &defaultPermMode,
						},
					},
				},
			},
		},
	}

	created, err := s.client.CoreV1().Pods(
		"default").Create(c, &expectedPod, metav1.CreateOptions{})
	if err != nil {
		c.JSON(500, err)
		return
	}
	expectedService := v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.IdentityToken,
			Namespace: "default",
			Labels: map[string]string{
				"name": req.IdentityToken,
			},
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"name": req.IdentityToken,
			},
			Type: v1.ServiceTypeClusterIP,
			Ports: []v1.ServicePort{
				{
					Name: "ssh",
					Port: 2222,
				},
			},
		},
	}
	_, err = s.client.CoreV1().Services("default").Create(c, &expectedService, metav1.CreateOptions{})
	if err != nil {
		c.JSON(500, err)
		return
	}

	resp := types.EnvironmentCreateResponse{
		ID: created.Name,
	}
	c.JSON(201, resp)
}
