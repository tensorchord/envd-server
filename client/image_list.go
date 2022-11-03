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

// ImageList lists the images.
func (cli *Client) ImageList(ctx context.Context, owner string) (types.ImageListResponse, error) {
	url := fmt.Sprintf("/users/%s/images", owner)
	resp, err := cli.get(ctx, url, nil, nil)
	defer ensureReaderClosed(resp)

	if err != nil {
		return types.ImageListResponse{}, wrapResponseError(err, resp, "owner", owner)
	}

	var response types.ImageListResponse
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}
