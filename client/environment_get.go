// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cockroachdb/errors"

	"github.com/tensorchord/envd-server/api/types"
)

// EnvironmentGet gets the environment.
func (cli *Client) EnvironmentGet(ctx context.Context, name string) (types.EnvironmentGetResponse, error) {
	username, headers, err := cli.getUserAndHeaders()
	if err != nil {
		return types.EnvironmentGetResponse{},
			errors.Wrap(err, "failed to get user and headers")
	}

	url := fmt.Sprintf("/users/%s/environments/%s", username, name)
	resp, err := cli.get(ctx, url, nil, headers)
	defer ensureReaderClosed(resp)

	if err != nil {
		return types.EnvironmentGetResponse{},
			wrapResponseError(err, resp, "user", username)
	}

	var response types.EnvironmentGetResponse
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}
