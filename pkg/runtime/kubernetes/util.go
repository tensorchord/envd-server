// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package kubernetes

import (
	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/consts"
	"github.com/tensorchord/envd-server/pkg/util"
	"github.com/tensorchord/envd-server/pkg/util/imageutil"
)

func extractResourceRequest(env types.Environment) (v1.ResourceList, error) {
	res := env.Resources
	resRequest := v1.ResourceList{}
	if res.CPU != "" {
		cpu, err := resource.ParseQuantity(res.CPU)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse cpu resource")
		}
		resRequest[v1.ResourceCPU] = cpu
	}
	if res.Memory != "" {
		mem, err := resource.ParseQuantity(res.Memory)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse memory resource")
		}
		resRequest[v1.ResourceMemory] = mem
	}
	if res.GPU != "" {
		gpu, err := resource.ParseQuantity(res.GPU)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse gpu resource")
		}
		resRequest[ResourceNvidiaGPU] = gpu
	}
	return resRequest, nil
}

func environmentFromKubernetesPod(pod *v1.Pod) (*types.Environment, error) {
	// Validate if the pod is an environment pod
	if pod == nil {
		return nil, errors.New("pod is nil")
	}

	env := &types.Environment{
		ObjectMeta: types.ObjectMeta{
			Name: pod.Name,
		},
		Spec:   types.EnvironmentSpec{},
		Status: types.EnvironmentStatus{},
	}

	// Convert the spec.
	if len(pod.Spec.Containers) != 0 {
		env.Spec.Image = pod.Spec.Containers[0].Image
		env.Spec.Env = envVarsfromKubernetes(pod.Spec.Containers[0].Env)
		env.Spec.Cmd = pod.Spec.Containers[0].Command
	}
	if pod.Labels[consts.PodLabelUID] != "" {
		env.Spec.Owner = pod.Labels[consts.PodLabelUID]
	}
	// Get the ports
	portLabel, ok := pod.Annotations[consts.ImageLabelPorts]
	if !ok {
		logrus.Info("failed to get port label")
		return nil, errors.New("failed to get port annotation")
	}
	ports, err := imageutil.PortsFromLabel(portLabel)
	if err != nil {
		logrus.Infof("failed to get ports from: %s", portLabel)
		return nil, errors.Wrap(err, "failed to get ports from label")
	}
	env.Spec.Ports = ports

	if jupyterAddr, ok := pod.Annotations[consts.PodLabelJupyterAddr]; ok {
		env.Status.JupyterAddr = &jupyterAddr
	}
	if rstudioServerAddr, ok := pod.Annotations[consts.PodLabelRStudioServerAddr]; ok {
		env.Status.RStudioServerAddr = &rstudioServerAddr
	}

	env.CreatedAt = pod.CreationTimestamp.Unix()

	// only reserve labels with prefix `ai.tensorchord.envd.`
	env.Labels = util.Filter(env.Labels, util.IsEnvdLabel)
	env.Status.Phase = string(pod.Status.Phase)
	return env, nil
}

func envVarsfromKubernetes(env []v1.EnvVar) []types.EnvVar {
	envVars := make([]types.EnvVar, len(env))
	for i, e := range env {
		envVars[i] = types.EnvVar{
			Name:  e.Name,
			Value: e.Value,
		}
	}
	return envVars
}
