// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/tensorchord/envd-server/api/types"
)

// ImageGet gets the image info.
func (cli *Client) ImageGet(ctx context.Context, owner, name string) (types.ImageGetResponse, error) {
	url := fmt.Sprintf("/users/%s/images/%s", owner, url.PathEscape(name))
	logrus.WithField("url", url).Debug("build image get url")
	resp, err := cli.get(ctx, url, nil, nil)
	defer ensureReaderClosed(resp)

	if err != nil {
		return types.ImageGetResponse{}, wrapResponseError(err, resp, "owner", owner)
	}

	var response types.ImageGetResponse
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}
