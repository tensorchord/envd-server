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

// EnvironmentList lists the environment.
func (cli *Client) EnvironmentList(ctx context.Context) (types.EnvironmentListResponse, error) {
	username, headers, err := cli.getUserAndHeaders()
	if err != nil {
		return types.EnvironmentListResponse{},
			errors.Wrap(err, "failed to get user and headers")
	}

	url := fmt.Sprintf("/users/%s/environments", username)
	resp, err := cli.get(ctx, url, nil, headers)
	defer ensureReaderClosed(resp)

	if err != nil {
		return types.EnvironmentListResponse{},
			wrapResponseError(err, resp, "username", username)
	}

	var response types.EnvironmentListResponse
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}
