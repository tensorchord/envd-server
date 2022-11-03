// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package util

import (
	"strings"

	"github.com/tensorchord/envd-server/pkg/consts"
)

// Filter filter elements in map
func Filter(origin map[string]string, predicate func(string) bool) map[string]string {
	result := make(map[string]string)
	for k, v := range origin {
		if predicate(k) {
			result[k] = v
		}
	}
	return result
}

// IsEnvdLabel check if envd label
func IsEnvdLabel(key string) bool {
	return strings.HasPrefix(key, consts.EnvdLabelPrefix)
}
