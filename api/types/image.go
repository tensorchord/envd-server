// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package types

import (
	dockertypes "github.com/docker/docker/api/types"
)

type ImageInfo struct {
	OwnerToken string
	Image      string
	Summary    dockertypes.ImageSummary
}

type ImageGetRequest struct {
	Name string `uri:"name" example:"pytorch-example"`
}

type ImageGetResponse struct {
	dockertypes.ImageSummary `json:",inline"`
}

type ImageListRequest struct {
}

type ImageListResponse struct {
	Items []ImageInfo `json:"items,omitempty"`
}
