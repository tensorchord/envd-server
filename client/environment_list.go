// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tensorchord/envd-server/api/types"
)

// EnvironmentList lists the environment.
func (cli *Client) EnvironmentList(ctx context.Context, req types.EnvironmentListRequest) (types.EnvironmentListResponse, error) {
	url := fmt.Sprintf("/users/%s/environments", req.IdentityToken)
	resp, err := cli.get(ctx, url, nil, nil)
	defer ensureReaderClosed(resp)

	if err != nil {
		return types.EnvironmentListResponse{}, err
	}

	var response types.EnvironmentListResponse
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}
