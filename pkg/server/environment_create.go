// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"context"
	"fmt"

	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/image"
	"github.com/gin-gonic/gin"
	imagespecv1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/consts"
	"github.com/tensorchord/envd-server/pkg/util/imageutil"
)

// @Summary     Create the environment.
// @Description Create the environment.
// @Tags        environment
// @Accept      json
// @Produce     json
// @Param       identity_token path     string                         true "identity token" example("a332139d39b89a241400013700e665a3")
// @Param       request        body     types.EnvironmentCreateRequest true "query params"
// @Success     200            {object} types.EnvironmentCreateResponse
// @Router      /users/{identity_token}/environments [post]
func (s *Server) environmentCreate(c *gin.Context) {
	it := c.GetString("identity_token")

	var req types.EnvironmentCreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(500, err)
		return
	}

	cfg, err := getImageConfig(c.Request.Context(), req.Spec.Image)
	if err != nil {
		c.JSON(500, err)
		return
	}
	// Merge image labels to pod.
	labels := map[string]string{
		consts.LabelUID:             it,
		consts.LabelEnvironmentName: req.Name,
	}

	logrus.WithFields(logrus.Fields{
		"identity_token": it,
		"image_labels":   cfg.Labels,
		"environment":    req.Environment,
	}).Debug("creating the environment")
	annotations := map[string]string{}
	for k, v := range cfg.Labels {
		annotations[k] = v
	}

	ports, err := imageutil.PortsFromLabel(cfg.Labels[consts.ImageLabelPorts])
	if err != nil {
		c.JSON(500, errors.Wrap(err, "failed to get ports from label"))
		return
	}
	v, ok := cfg.Labels[consts.ImageLabelRepo]
	repoInfo := &types.EnvironmentRepoInfo{}
	if ok {
		repoInfo, err = imageutil.RepoInfoFromLabel(v)
		if err != nil {
			c.JSON(500, errors.Wrap(err, "failed to get repo information from label"))
			return
		}
	}
	workDir := cfg.Labels[consts.ContainerLabelName]
	hostKeyPath := "/var/envd/hostkey"
	authKeyPath := "/var/envd/authkey"
	var defaultPermMode int32 = 0666
	expectedPod := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        req.Name,
			Namespace:   "default",
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "envd",
					Image: req.Spec.Image,
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
							SecretName:  "envd-server",
							DefaultMode: &defaultPermMode,
						},
					},
				},
			},
		},
	}
	if len(repoInfo.URL) > 0 {
		logrus.Debugf("clone code from %s", repoInfo.URL)
		expectedPod.Spec.InitContainers = append(expectedPod.Spec.InitContainers, v1.Container{
			Name:  "git-cloner",
			Image: "alpine/git",
			Args:  []string{"clone", "--", repoInfo.URL, "/code"},
			VolumeMounts: []v1.VolumeMount{
				{
					Name:      "code-dir",
					MountPath: "/code",
				},
			},
		})
		expectedPod.Spec.Containers[0].VolumeMounts = append(expectedPod.Spec.Containers[0].VolumeMounts, v1.VolumeMount{
			Name:      "code-dir",
			MountPath: fmt.Sprintf("/home/envd/%s", workDir),
		})
		expectedPod.Spec.Volumes = append(expectedPod.Spec.Volumes, v1.Volume{
			Name: "code-dir",
			VolumeSource: v1.VolumeSource{
				HostPath: &v1.HostPathVolumeSource{
					Path: fmt.Sprintf("/var/envd/code/%s", req.Name),
				},
			},
		})
	}

	_, err = s.client.CoreV1().Pods(
		"default").Create(c, &expectedPod, metav1.CreateOptions{})
	if err != nil {
		logrus.Infof("failed to create pod: %v", err)
		c.JSON(500, err)
		return
	}

	expectedService := v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.Name,
			Namespace: "default",
			Labels:    labels,
		},
		Spec: v1.ServiceSpec{
			Selector: labels,
			Type:     v1.ServiceTypeClusterIP,
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
		Created: req.Environment,
	}
	resp.Created.Spec.Ports = ports
	c.JSON(201, resp)
}

func getImageConfig(ctx context.Context, imagename string) (
	*imagespecv1.ImageConfig, error) {
	ref, err := docker.ParseReference(fmt.Sprintf("//%s", imagename))
	if err != nil {
		return nil, err
	}
	src, err := ref.NewImageSource(ctx, nil)
	if err != nil {
		return nil, err
	}
	img, err := image.FromUnparsedImage(ctx, nil, image.UnparsedInstance(src, nil))
	if err != nil {
		return nil, err
	}
	c, err := img.OCIConfig(ctx)
	if err != nil {
		return nil, err
	}
	return &c.Config, nil
}
