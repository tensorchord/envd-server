// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package types

// AuthNRequest contains authorization information for connecting to a envd server.
type AuthNRequest struct {
	// LoginName is used to authenticate the user and get
	// an access token for the registry.
	LoginName string `json:"login_name,omitempty" example:"alice"`

	// Password stores the hashed password.
	Password string `json:"password,omitempty"`
}

type AuthStatus string

const (
	AuthSuccess AuthStatus = "Login succeeded"
	AuthFail    AuthStatus = "Login Fail"
	Error       AuthStatus = "Internal_Error"
)

type AuthNResponse struct {
	// LoginName is used to authenticate the user and get
	// an access token for the registry.
	LoginName string `json:"login_name,omitempty" example:"alice"`
	// An opaque token used to authenticate a user after a successful login
	// Required: true
	IdentityToken string `json:"identity_token" example:"a332139d39b89a241400013700e665a3"`
	// The status of the authentication
	// Required: true
	Status AuthStatus `json:"status" example:"Login successfully"`
}
