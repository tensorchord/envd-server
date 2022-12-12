// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package kubernetes

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/errdefs"
	"github.com/tensorchord/envd-server/pkg/consts"
)

func (p generalProvisioner) EnvironmentGet(ctx context.Context,
	owner, envName string) (*types.Environment, error) {

	pod, err := p.client.CoreV1().
		Pods(p.namespace).Get(ctx, envName, metav1.GetOptions{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			return nil, errdefs.NotFound(err)
		}
		return nil, errors.Wrap(err, "failed to get pod")
	}
	if pod.Labels[consts.PodLabelUID] != owner {
		logrus.WithFields(logrus.Fields{
			"loginname_in_pod":     pod.Labels[consts.PodLabelUID],
			"loginname_in_request": owner,
		}).Debug("mismatch loginname")
		return nil, errdefs.Unauthorized(errors.New("mismatch loginname"))
	}

	if pod == nil {
		return nil, errdefs.NotFound(errors.New("pod not found"))
	}

	e, err := environmentFromKubernetesPod(pod)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert pod to environment")
	}
	return e, nil
}
