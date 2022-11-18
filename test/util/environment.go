// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package util

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/tensorchord/envd-server/pkg/consts"
)

func NewPod(name, it string) *v1.Pod {
	var defaultPermMode int32 = 0666

	p := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
			Labels: map[string]string{
				consts.PodLabelUID: it,
			},
			Annotations: map[string]string{
				consts.ImageLabelPorts: `
[{"name": "ssh", "port": 2222}]`,
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "envd",
					Image: "test",
					Ports: []v1.ContainerPort{
						{
							Name:          "ssh",
							ContainerPort: 2222,
						},
					},
					Env: []v1.EnvVar{
						{
							Name:  "ENVD_HOST_KEY",
							Value: "test",
						},
						{
							Name:  "ENVD_AUTHORIZED_KEYS_PATH",
							Value: "test",
						},
						{
							Name:  "ENVD_WORKDIR",
							Value: fmt.Sprintf("/home/envd/%s", "test"),
						},
					},
					VolumeMounts: []v1.VolumeMount{
						{
							Name:      "secret",
							ReadOnly:  true,
							MountPath: "test",
							SubPath:   "hostkey",
						},
						{
							Name:      "secret",
							ReadOnly:  true,
							MountPath: "test",
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
	return &p
}
