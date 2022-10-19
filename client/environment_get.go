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
func (cli *Client) EnvironmentGet(ctx context.Context, owner, name string) (types.EnvironmentGetResponse, error) {
	url := fmt.Sprintf("/users/%s/environments/%s", owner, name)
	resp, err := cli.get(ctx, url, nil, nil)
	defer ensureReaderClosed(resp)

	if err != nil {
		return types.EnvironmentGetResponse{}, wrapResponseError(err, resp, "owner", owner)
	}

	var response types.EnvironmentGetResponse
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}
