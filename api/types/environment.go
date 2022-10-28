// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package types

type Environment struct {
	ObjectMeta `json:",inline"`
	Spec       EnvironmentSpec   `json:"spec,omitempty"`
	Status     EnvironmentStatus `json:"status,omitempty"`
}

type ObjectMeta struct {
	Name string `json:"name,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`
}

type EnvironmentSpec struct {
	Owner string            `json:"owner,omitempty"`
	Image string            `json:"image,omitempty"`
	Env   []string          `json:"env,omitempty"`
	Cmd   []string          `json:"cmd,omitempty"`
	Ports []EnvironmentPort `json:"ports,omitempty"`
	// TODO(gaocegege): Add volume specific spec.
}

type EnvironmentStatus struct {
	Phase string `json:"phase,omitempty"`
}

type EnvironmentPort struct {
	Name string `json:"name,omitempty"`
	Port int32  `json:"port,omitempty"`
}

type EnvironmentRepoInfo struct {
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

type EnvironmentCreateRequest struct {
	Environment `json:",inline,omitempty"`
}

type EnvironmentCreateResponse struct {
	Created Environment `json:"environment,omitempty"`

	// Warnings encountered when creating the pod
	// Required: true
	Warnings []string `json:"warnings,omitempty"`
}

type EnvironmentListRequest struct {
}

type EnvironmentListResponse struct {
	Items []Environment `json:"items,omitempty"`
}

type EnvironmentRemoveRequest struct {
	Name string `uri:"name" example:"pytorch-example"`
}

type EnvironmentRemoveResponse struct {
}

type EnvironmentGetRequest struct {
	Name string `uri:"name" example:"pytorch-example"`
}

type EnvironmentGetResponse struct {
	Environment `json:",inline"`
}
