// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package types

type Environment struct {
	Spec   EnvironmentSpec
	Status EnvironmentStatus
}

type EnvironmentSpec struct {
	Name  string            `json:"name,omitempty"`
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

type EnvironmentCreateRequest struct {
	EnvironmentSpec `json:",inline"`
}

type EnvironmentCreateResponse struct {
	// The ID of the created container
	// Required: true
	ID string `json:"id"`

	// Warnings encountered when creating the pod
	// Required: true
	Warnings []string `json:"warnings"`
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
