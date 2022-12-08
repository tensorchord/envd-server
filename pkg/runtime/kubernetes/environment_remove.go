// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package kubernetes

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/tensorchord/envd-server/errdefs"
	"github.com/tensorchord/envd-server/pkg/consts"
)

func (p generalProvisioner) EnvironmentRemove(ctx context.Context,
	owner, envName string) error {
	logger := logrus.WithField("env", envName)
	pod, err := p.client.CoreV1().Pods(p.namespace).
		Get(ctx, envName, metav1.GetOptions{})
	if !k8serrors.IsNotFound(err) {
		if err != nil {
			return errors.Wrap(err, "failed to get pod")
		}
		if pod.Labels[consts.PodLabelUID] != owner {
			logger.WithFields(logrus.Fields{
				"loginname_in_pod":     pod.Labels[consts.PodLabelUID],
				"loginname_in_request": owner,
			}).Debug("mismatch loginname")
			return errdefs.Unauthorized(errors.New("mismatch loginname"))
		}

		err = p.client.CoreV1().Pods(
			p.namespace).Delete(ctx, envName, metav1.DeleteOptions{})
		if err != nil && !k8serrors.IsNotFound(err) {
			return errors.Wrap(err, "failed to delete pod")
		}
		logger.Debugf("pod %s is deleted", envName)
	}

	service, err := p.client.CoreV1().Services(p.namespace).
		Get(ctx, envName, metav1.GetOptions{})
	if !k8serrors.IsNotFound(err) {
		if err != nil {
			return errors.Wrap(err, "failed to get service")
		}
		if service.Labels[consts.PodLabelUID] != owner {
			logger.WithFields(logrus.Fields{
				"loginname_in_pod":     pod.Labels[consts.PodLabelUID],
				"loginname_in_request": owner,
			}).Debug("mismatch loginname")
			return errdefs.Unauthorized(errors.New("mismatch loginname"))
		}
		err = p.client.CoreV1().Services(p.namespace).
			Delete(ctx, envName, metav1.DeleteOptions{})
		if err != nil && !k8serrors.IsNotFound(err) {
			return errors.Wrap(err, "failed to delete service")
		}
		logger.Debugf("service %s is deleted", envName)
	}

	return nil
}
