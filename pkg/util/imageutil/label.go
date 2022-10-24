// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package imageutil

import (
	"encoding/json"

	"github.com/tensorchord/envd-server/api/types"
)

func PortsFromLabel(label string) ([]types.EnvironmentPort, error) {
	var ports []types.EnvironmentPort
	if err := json.Unmarshal([]byte(label), &ports); err != nil {
		return nil, err
	}

	return ports, nil
}
