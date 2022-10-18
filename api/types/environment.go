// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package types

import v1 "k8s.io/api/core/v1"

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
	// TODO(gaocegege): Move it to URI.
	// Use auth instead of in the request body.
	IdentityToken string `json:"identity_token"  example:"a332139d39b89a241400013700e665a3"`
	Image         string `json:"image,omitempty"`
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
	// Use auth instead of in the request body.
	IdentityToken string `uri:"identity_token"  example:"a332139d39b89a241400013700e665a3"`
}

type EnvironmentListResponse struct {
	Pod v1.Pod
}

type EnvironmentRemoveRequest struct {
	// Use auth instead of in the requrest body.
	IdentityToken string `uri:"identity_token"  example:"a332139d39b89a241400013700e665a3"`
}

type EnvironmentRemoveResponse struct {
}
