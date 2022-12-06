// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/cockroachdb/errors"
	"github.com/tensorchord/envd-server/api/types"
)

// EnvironmentCreate creates the environment.
func (cli *Client) EnvironmentCreate(ctx context.Context,
	req types.EnvironmentCreateRequest) (types.EnvironmentCreateResponse, error) {

	username, headers, err := cli.getUserAndHeaders()
	if err != nil {
		return types.EnvironmentCreateResponse{},
			errors.Wrap(err, "failed to get user and headers")
	}

	urlString := fmt.Sprintf("/users/%s/environments", username)
	resp, err := cli.post(ctx, urlString, url.Values{}, req, headers)
	defer ensureReaderClosed(resp)

	if err != nil {
		return types.EnvironmentCreateResponse{}, err
	}

	var response types.EnvironmentCreateResponse
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}
