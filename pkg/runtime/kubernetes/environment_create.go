// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package kubernetes

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/tensorchord/envd-server/api/types"
	servertypes "github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/consts"
)

func (p generalProvisioner) EnvironmentCreate(ctx context.Context,
	owner string, env servertypes.Environment,
	meta types.ImageMeta) (*servertypes.Environment, error) {
	resRequest, err := extractResourceRequest(env)
	if err != nil {
		return nil, errors.Wrap(err, "failed to extract resource request")
	}

	labels := map[string]string{
		consts.PodLabelUID:             owner,
		consts.PodLabelEnvironmentName: env.Name,
	}

	logrus.WithFields(logrus.Fields{
		"login-name":   owner,
		"image_labels": meta.Labels,
		"environment":  env,
	}).Debug("prepare to create the environment")
	annotations := map[string]string{}
	for k, v := range meta.Labels {
		annotations[k] = v
	}

	repoLabel, ok := meta.Labels[consts.ImageLabelRepo]
	repoInfo := &types.EnvironmentRepoInfo{}
	if ok {
		repoInfo, err = repoInfoFromLabel(repoLabel)
		if err != nil {
			logrus.Info("failed to parse repo from label")
			return nil, errors.Wrap(err, "failed to parse repo from label")
		}
	}

	projectName, ok := meta.Labels[consts.ImageLabelContainerName]
	if !ok {
		logrus.Info("failed to get the project name from label")
		return nil, errors.Wrap(err, "failed to get the project name from label")
	}
	logrus.WithFields(logrus.Fields{
		"repo":    repoInfo,
		"project": projectName,
	}).Debug("creating environment")
	hostKeyPath := "/var/envd/hostkey"
	authKeyPath := "/var/envd/authkey"
	var defaultPermMode int32 = 0666

	codeDirectoryVolumeMount := v1.VolumeMount{
		Name:      "code-dir",
		MountPath: fmt.Sprintf("/home/envd/%s", projectName),
	}

	expectedPod := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        env.Name,
			Namespace:   p.namespace,
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "envd",
					Image: env.Spec.Image,
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
						codeDirectoryVolumeMount,
					},
					Resources: v1.ResourceRequirements{
						Requests: resRequest,
						Limits:   resRequest,
					},
				},
				{
					Name:  "syncthing",
					Image: "linuxserver/syncthing:1.22.2",
					Ports: []v1.ContainerPort{
						{
							Name:          "syncthing",
							ContainerPort: 8384,
						},
						{
							Name:          "st-listen",
							Protocol:      v1.ProtocolTCP,
							ContainerPort: 22000,
						},
						{
							Name:          "st-discover",
							Protocol:      v1.ProtocolUDP,
							ContainerPort: 22000,
						},
					},
					VolumeMounts: []v1.VolumeMount{
						codeDirectoryVolumeMount,
						{
							Name:      "syncthing-config",
							MountPath: "/config",
						},
					},
					// TODO: define resource limit
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
				{
					Name: "code-dir",
					VolumeSource: v1.VolumeSource{
						EmptyDir: &v1.EmptyDirVolumeSource{},
						// HostPath: &v1.HostPathVolumeSource{
						// 	Path: fmt.Sprintf("/var/envd/code/%s", req.Name),
						// },
					},
				},
				{
					Name: "syncthing-config",
					VolumeSource: v1.VolumeSource{
						EmptyDir: &v1.EmptyDirVolumeSource{},
					},
				},
			},
		},
	}

	if repoInfo != nil && len(repoInfo.URL) > 0 {
		// TODO: Figure out how the clone directory works with the sync
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

	}

	if p.imagePullSecretName != nil {
		expectedPod.Spec.ImagePullSecrets = []v1.LocalObjectReference{
			{
				Name: *p.imagePullSecretName,
			},
		}
	}

	created, err := p.client.CoreV1().Pods(
		p.namespace).Create(ctx, &expectedPod, metav1.CreateOptions{})
	if err != nil {
		logrus.Infof("failed to create pod: %v", err)
		return nil, errors.Wrap(err, "failed to create pod")
	}

	expectedService := v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      env.Name,
			Namespace: p.namespace,
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
				{
					Name:       "syncthing",
					Port:       8384,
					TargetPort: intstr.FromInt(8384),
				},
				{
					Name:       "st-listen",
					Port:       22000,
					TargetPort: intstr.FromInt(22000),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "st-discover",
					Port:       22000,
					TargetPort: intstr.FromInt(22000),
					Protocol:   v1.ProtocolUDP,
				},
			},
		},
	}
	_, err = p.client.CoreV1().
		Services(p.namespace).Create(ctx, &expectedService, metav1.CreateOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create service")
	}

	createdEnv, err := environmentFromKubernetesPod(created)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create environment")
	}
	return createdEnv, nil
}
