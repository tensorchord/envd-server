// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package types

import v1 "k8s.io/api/core/v1"

type EnvironmentCreateRequest struct {
	// Use auth instead of in the requrest body.
	IdentityToken string `json:"identity_token,omitempty"`
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
	// Use auth instead of in the requrest body.
	IdentityToken string `json:"identity_token,omitempty" uri:"identity_token"`
}

type EnvironmentListResponse struct {
	Pod v1.Pod
}
