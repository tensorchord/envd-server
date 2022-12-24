// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package types

type ImageInfo struct {
	OwnerToken string
	ImageMeta
}

type ImageMeta struct {
	Name           string            `json:"name,omitempty" example:"pytorch-cuda:dev"`
	Digest         string            `json:"digest,omitempty"`
	Created        int64             `json:"created,omitempty"`
	Size           int64             `json:"size,omitempty"`
	Labels         map[string]string `json:"labels,omitempty"`
	APTPackages    []string          `json:"apt_packages,omitempty"`
	PythonCommands []string          `json:"python_commands,omitempty"`
	Ports          []EnvironmentPort `json:"ports,omitempty"`
}

type ImageGetRequest struct {
	Name string `uri:"name" example:"pytorch-example"`
}

type ImageGetResponse struct {
	ImageMeta `json:",inline"`
}

type ImageListRequest struct {
}

type ImageListResponse struct {
	Items []ImageMeta `json:"items,omitempty"`
}
