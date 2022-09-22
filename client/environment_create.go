package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/tensorchord/envd-server/api/types"
)

// EnvironmentCreate creates the environment.
func (cli *Client) EnvironmentCreate(ctx context.Context, auth types.EnvironmentCreateRequest) (types.EnvironmentCreateResponse, error) {
	resp, err := cli.post(ctx, "/environments", url.Values{}, auth, nil)
	defer ensureReaderClosed(resp)

	if err != nil {
		return types.EnvironmentCreateResponse{}, err
	}

	var response types.EnvironmentCreateResponse
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}
