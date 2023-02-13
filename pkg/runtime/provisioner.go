// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package runtime

import (
	"context"

	"github.com/tensorchord/envd-server/api/types"
)

type Provisioner interface {
	EnvironmentCreate(ctx context.Context,
		owner string, env types.Environment,
		meta types.ImageMeta) (*types.Environment, error)
	EnvironmentGet(ctx context.Context,
		owner, envName string) (*types.Environment, error)
	EnvironmentRemove(ctx context.Context,
		owner, envName string) error
	EnvironmentList(ctx context.Context,
		username string) ([]types.Environment, error)
}
