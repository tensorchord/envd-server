// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package client

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
)

// EnvironmentRemove the environment.
func (cli *Client) EnvironmentRemove(ctx context.Context,
	name string) error {
	username, headers, err := cli.getUserAndHeaders()
	if err != nil {
		return errors.Wrap(err, "failed to get user and headers")
	}

	url := fmt.Sprintf("/users/%s/environments/%s", username, name)
	resp, err := cli.delete(ctx, url, nil, headers)
	defer ensureReaderClosed(resp)
	return wrapResponseError(err, resp, "environment", name)
}
