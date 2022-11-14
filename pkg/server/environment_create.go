// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/image"
	containertypes "github.com/containers/image/v5/types"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgtype"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/consts"
	"github.com/tensorchord/envd-server/pkg/query"
	"github.com/tensorchord/envd-server/pkg/util/imageutil"
)

// @Summary     Create the environment.
// @Description Create the environment.
// @Tags        environment
// @Accept      json
// @Produce     json
// @Param       identity_token path     string                         true "identity token" example("a332139d39b89a241400013700e665a3")
// @Param       request        body     types.EnvironmentCreateRequest true "query params"
// @Success     201            {object} types.EnvironmentCreateResponse
// @Router      /users/{identity_token}/environments [post]
func (s *Server) environmentCreate(c *gin.Context) {
	it := c.GetString("identity_token")

	var req types.EnvironmentCreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(500, err)
		return
	}

	meta, err := getImageMeta(c.Request.Context(), req.Spec.Image)
	if err != nil {
		c.JSON(500, err)
		return
	}
	var pglabel pgtype.JSONB
	err = pglabel.Set(meta.Labels)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	_, err = s.Queries.CreateImageInfo(context.Background(),
		query.CreateImageInfoParams{OwnerToken: it,
			Name: meta.Name, Digest: meta.Digest, Created: meta.Created, Size: meta.Size, Labels: pglabel})
	if err != nil {
		c.JSON(500, err)
	}
	labels := map[string]string{
		consts.PodLabelUID:             it,
		consts.PodLabelEnvironmentName: req.Name,
	}

	logrus.WithFields(logrus.Fields{
		"identity_token": it,
		"image_labels":   meta.Labels,
		"environment":    req.Environment,
	}).Debug("prepare to create the environment")
	annotations := map[string]string{}
	for k, v := range meta.Labels {
		annotations[k] = v
	}

	portLabel, ok := meta.Labels[consts.ImageLabelPorts]
	if !ok {
		logrus.Info("failed to get port label")
		c.JSON(500, errors.Wrap(err, "failed to get the port"))
		return
	}
	ports, err := imageutil.PortsFromLabel(portLabel)
	if err != nil {
		logrus.Infof("failed to get ports from: %s", portLabel)
		c.JSON(500, errors.Wrap(err, "failed to parse ports from label"))
		return
	}

	repoLabel, ok := meta.Labels[consts.ImageLabelRepo]
	repoInfo := &types.EnvironmentRepoInfo{}
	if ok {
		repoInfo, err = imageutil.RepoInfoFromLabel(repoLabel)
		if err != nil {
			logrus.Info("failed to parse repo from label")
			c.JSON(500, errors.Wrap(err, "failed to get repo information from label"))
			return
		}
	}

	projectName, ok := meta.Labels[consts.ImageLabelContainerName]
	if !ok {
		logrus.Info("failed to get the project name from label")
		c.JSON(500, errors.New("failed to get the project name(working dir) from label"))
		return
	}
	logrus.WithFields(logrus.Fields{
		"port":    ports,
		"repo":    repoInfo,
		"project": projectName,
	}).Debug("creating environment")
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
						{
							Name:  "ENVD_WORKDIR",
							Value: fmt.Sprintf("/home/envd/%s", projectName),
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
	if repoInfo != nil && len(repoInfo.URL) > 0 {
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
			MountPath: fmt.Sprintf("/home/envd/%s", projectName),
		})
		expectedPod.Spec.Volumes = append(expectedPod.Spec.Volumes, v1.Volume{
			Name: "code-dir",
			VolumeSource: v1.VolumeSource{
				EmptyDir: &v1.EmptyDirVolumeSource{},
				// HostPath: &v1.HostPathVolumeSource{
				// 	Path: fmt.Sprintf("/var/envd/code/%s", req.Name),
				// },
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

func getImageMeta(ctx context.Context, imageName string) (
	meta types.ImageMeta, err error) {
	ref, err := docker.ParseReference(fmt.Sprintf("//%s", imageName))
	if err != nil {
		return
	}
	sys := &containertypes.SystemContext{}
	src, err := ref.NewImageSource(ctx, sys)
	if err != nil {
		return
	}
	digest, err := docker.GetDigest(ctx, sys, ref)
	if err != nil {
		return
	}
	image, err := image.FromUnparsedImage(ctx, sys, image.UnparsedInstance(src, &digest))
	if err != nil {
		return
	}
	inspect, err := image.Inspect(ctx)
	if err != nil {
		return
	}
	size, err := image.Size()
	if err != nil {
		return
	}
	if size < 0 {
		// correct the image size
		size = 0
		for _, layer := range inspect.LayersData {
			size += layer.Size
		}
	}
	meta = types.ImageMeta{
		Name:    imageName,
		Created: inspect.Created.Unix(),
		Digest:  string(digest),
		Labels:  inspect.Labels,
		Size:    size,
	}
	logrus.WithField("image meta", meta).Debug("get image meta before creating env")
	src.Close()
	return meta, nil
}
