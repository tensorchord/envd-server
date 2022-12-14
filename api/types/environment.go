// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package types

type Environment struct {
	ObjectMeta `json:",inline"`
	Spec       EnvironmentSpec   `json:"spec,omitempty"`
	Status     EnvironmentStatus `json:"status,omitempty"`
	Resources  ResourceSpec      `json:"resource,omitempty"`
	CreatedAt  int64             `json:"created_at,omitempty"`
}

type ResourceSpec struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
	GPU    string `json:"gpu,omitempty"`
}

type ObjectMeta struct {
	Name string `json:"name,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`
}

type EnvironmentSpec struct {
	Owner string            `json:"owner,omitempty"`
	Image string            `json:"image,omitempty"`
	Env   []EnvVar          `json:"env,omitempty"`
	Cmd   []string          `json:"cmd,omitempty"`
	Ports []EnvironmentPort `json:"ports,omitempty"`
	// TODO(gaocegege): Add volume specific spec.
}

type EnvVar struct {
	// Name of the environment variable. Must be a C_IDENTIFIER.
	Name string `json:"name"`

	// Optional: no more than one of the following may be specified.

	// Variable references $(VAR_NAME) are expanded
	// using the previously defined environment variables in the container and
	// any service environment variables. If a variable cannot be resolved,
	// the reference in the input string will be unchanged. Double $$ are reduced
	// to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.
	// "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)".
	// Escaped references will never be expanded, regardless of whether the variable
	// exists or not.
	// Defaults to "".
	// +optional
	Value string `json:"value,omitempty"`
}

type EnvironmentStatus struct {
	Phase             string  `json:"phase,omitempty"`
	JupyterAddr       *string `json:"jupyter_addr,omitempty"`
	RStudioServerAddr *string `json:"rstudio_server_addr,omitempty"`
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
