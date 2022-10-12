// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package types

// AuthRequest contains authorization information for connecting to a envd server.
type AuthRequest struct {
	// Username  string `json:"username,omitempty"`
	// Password  string `json:"password,omitempty"`
	PublicKey string `json:"public_key"`

	// IdentityToken is used to authenticate the user and get
	// an access token for the registry.
	// Required: true
	IdentityToken string `json:"identity_token" example:"a332139d39b89a241400013700e665a3"`
}

type AuthResponse struct {
	// An opaque token used to authenticate a user after a successful login
	// Required: true
	IdentityToken string `json:"identity_token" example:"a332139d39b89a241400013700e665a3"`
	// The status of the authentication
	// Required: true
	Status string `json:"status" example:"Login successfully"`
}
