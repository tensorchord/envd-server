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
	"github.com/sirupsen/logrus"
	"github.com/tensorchord/envd-server/api/types"
)

// ImageGet gets the image info.
func (cli *Client) ImageGet(
	ctx context.Context, name string) (types.ImageGetResponse, error) {
	username, headers, err := cli.getUserAndHeaders()
	if err != nil {
		return types.ImageGetResponse{},
			errors.Wrap(err, "failed to get user and headers")
	}

	url := fmt.Sprintf("/users/%s/images/%s", username, url.PathEscape(name))
	logrus.WithField("url", url).Debug("build image get url")
	resp, err := cli.get(ctx, url, nil, headers)
	defer ensureReaderClosed(resp)

	if err != nil {
		return types.ImageGetResponse{},
			wrapResponseError(err, resp, "username", username)
	}

	var response types.ImageGetResponse
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}
