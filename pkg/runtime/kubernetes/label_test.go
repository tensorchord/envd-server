// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package kubernetes

import (
	"reflect"
	"testing"

	"github.com/tensorchord/envd-server/api/types"
)

func TestPortsFromLabel(t *testing.T) {
	tcs := []struct {
		label       string
		expectedErr bool
		port        []types.EnvironmentPort
	}{
		{
			label:       `[{"name": "test", "port": 2222}]`,
			expectedErr: false,
			port: []types.EnvironmentPort{
				{
					Name: "test",
					Port: 2222,
				},
			},
		},
		{
			label:       ``,
			expectedErr: false,
			port:        nil,
		},
		{
			label:       `[{"name": "test", "port": 2222},{"name": "jupyter", "port": 8080}]`,
			expectedErr: false,
			port: []types.EnvironmentPort{
				{
					Name: "test",
					Port: 2222,
				},
				{
					Name: "jupyter",
					Port: 8080,
				},
			},
		},
	}

	for _, tc := range tcs {
		p, err := portsFromLabel(tc.label)
		if tc.expectedErr {
			if err == nil {
				t.Errorf("Expected err, got nil")
			}
			continue
		}

		if e := reflect.DeepEqual(tc.port, p); e != true {
			t.Errorf("Expected ports %v, got %v", tc.port, p)
		}
	}
}
