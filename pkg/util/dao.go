// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package util

import (
	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/query"
)

func DaoToImageMeta(dao query.ImageInfo) (*types.ImageMeta, error) {
	var label map[string]string
	if err := dao.Labels.AssignTo(&label); err != nil {
		return nil, err
	}
	meta := types.ImageMeta{
		Name:    dao.Name,
		Digest:  dao.Digest,
		Created: dao.Created,
		Size:    dao.Size,
		Labels:  label,
	}
	return &meta, nil
}
