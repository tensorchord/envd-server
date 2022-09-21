package types

// AuthRequest contains authorization information for connecting to a envd server.
type AuthRequest struct {
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	PublicKey string `json:"public_key,omitempty"`

	// IdentityToken is used to authenticate the user and get
	// an access token for the registry.
	IdentityToken string `json:"identity_token,omitempty"`
}

type AuthResponse struct {
	// An opaque token used to authenticate a user after a successful login
	// Required: true
	IdentityToken string `json:"identity_token"`
	// The status of the authentication
	// Required: true
	Status string `json:"status"`
}
