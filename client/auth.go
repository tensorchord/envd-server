// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

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
