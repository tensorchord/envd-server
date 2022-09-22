package types

type EnvironmentCreateRequest struct {
	// Use auth instead of in the requrest body.
	IdentityToken string `json:"identity_token,omitempty"`
	Image         string `json:"image,omitempty"`
}

type EnvironmentCreateResponse struct {
	// The ID of the created container
	// Required: true
	ID string `json:"Id"`

	// Warnings encountered when creating the pod
	// Required: true
	Warnings []string `json:"Warnings"`
}
