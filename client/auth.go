package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/tensorchord/envd-server/api/types"
)

// Auth authenticates the envd server.
// It returns unauthorizedError when the authentication fails.
func (cli *Client) Auth(ctx context.Context, auth types.AuthRequest) (types.AuthResponse, error) {
	resp, err := cli.post(ctx, "/auth", url.Values{}, auth, nil)
	defer ensureReaderClosed(resp)

	if err != nil {
		return types.AuthResponse{}, err
	}

	var response types.AuthResponse
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}
