// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package client

import (
	"context"
	"fmt"

	"github.com/tensorchord/envd-server/api/types"
)

// EnvironmentList lists the environment.
func (cli *Client) EnvironmentRemove(ctx context.Context,
	req types.EnvironmentRemoveRequest) error {
	url := fmt.Sprintf("/users/%s/environments", req.IdentityToken)
	resp, err := cli.delete(ctx, url, nil, nil)
	defer ensureReaderClosed(resp)
	return wrapResponseError(err, resp, "environment", req.IdentityToken)
}
