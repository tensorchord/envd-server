// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package kubernetes

import (
	"context"

	"github.com/pkg/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/errdefs"
	"github.com/tensorchord/envd-server/pkg/consts"
)

func (p generalProvisioner) EnvironmentList(ctx context.Context,
	owner string) ([]types.Environment, error) {
	ls := labels.Set{
		consts.PodLabelUID: owner,
	}

	pods, err := p.client.CoreV1().Pods(
		p.namespace).List(ctx, metav1.ListOptions{
		LabelSelector: ls.String(),
	})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			return nil, errdefs.NotFound(err)
		}
		return nil, errors.Wrap(err, "failed to get pods")
	}

	res := []types.Environment{}

	for _, pod := range pods.Items {
		e, err := environmentFromKubernetesPod(&pod)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert pod to environment")
		}
		res = append(res, *e)
	}

	return res, nil
}
