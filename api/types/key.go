// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package types

type KeyCreateRequest struct {
	// Name is the key name.
	Name string `json:"name,omitempty" example:"mykey"`
	// PublicKey is the ssh public key
	PublicKey string `json:"public_key,omitempty"`
}

type KeyCreateResponse struct {
	// Name is the key name.
	Name string `json:"name,omitempty" example:"mykey"`
	// PublicKey is the ssh public key
	PublicKey string `json:"public_key,omitempty"`
	// LoginName is used to authenticate the user and get
	// an access token for the registry.
	LoginName string `json:"login_name,omitempty" example:"alice"`
}
