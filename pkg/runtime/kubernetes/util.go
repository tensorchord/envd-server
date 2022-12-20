// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package kubernetes

import (
	"strings"

	"github.com/cockroachdb/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/consts"
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
		return nil, errors.New("failed to get port annotation")
	}
	ports, err := portsFromLabel(portLabel)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ports from label")
	}
	env.Spec.Ports = ports

	// Get the apt packages
	if aptLabel, ok := pod.Annotations[consts.ImageLabelAPTPackages]; ok {
		packages, err := aptPackagesFromLabel(aptLabel)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get apt packages from label")
		}
		env.Spec.APTPackages = packages
	}

	// Get the python commands
	if pythonLabel, ok := pod.Annotations[consts.ImageLabelPythonCommands]; ok {
		pythonCommands, err := pythonCommandsFromLabel(pythonLabel)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get python commands from label")
		}
		env.Spec.PythonCommands = pythonCommands
	}

	if jupyterAddr, ok := pod.Annotations[consts.PodLabelJupyterAddr]; ok {
		env.Status.JupyterAddr = &jupyterAddr
	}
	if rstudioServerAddr, ok := pod.Annotations[consts.PodLabelRStudioServerAddr]; ok {
		env.Status.RStudioServerAddr = &rstudioServerAddr
	}

	env.CreatedAt = pod.CreationTimestamp.Unix()

	// only reserve labels with prefix `ai.tensorchord.envd.`
	env.Labels = filter(pod.Labels, isEnvdLabel)
	env.Annotations = filter(pod.Annotations, isEnvdLabel)
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

// Filter filter elements in map
func filter(origin map[string]string, predicate func(string) bool) map[string]string {
	result := make(map[string]string)
	for k, v := range origin {
		if predicate(k) {
			result[k] = v
		}
	}
	return result
}

// IsEnvdLabel check if envd label
func isEnvdLabel(key string) bool {
	return strings.HasPrefix(key, consts.EnvdLabelPrefix)
}
