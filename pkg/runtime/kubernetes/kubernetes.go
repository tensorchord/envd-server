// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package kubernetes

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"

	"github.com/tensorchord/envd-server/pkg/runtime"
)

type generalProvisioner struct {
	client kubernetes.Interface
	logger *logrus.Entry

	namespace           string
	imagePullSecretName *string
}

func NewProvisioner(client kubernetes.Interface,
	namespace, imagePullSecretName string) runtime.Provisioner {
	p := &generalProvisioner{
		client:    client,
		namespace: namespace,
		logger: logrus.WithFields(logrus.Fields{
			"namespace":              namespace,
			"image-pull-secret-name": imagePullSecretName,
		}),
	}
	if imagePullSecretName != "" {
		p.imagePullSecretName = &imagePullSecretName
	}
	return p
}
