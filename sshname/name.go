// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package sshname

import (
	"fmt"
	"strings"
)

func Username(owner, envName string) (string, error) {
	return fmt.Sprintf("%s/%s", owner, envName), nil
}

func GetInfo(username string) (string, string, error) {
	s := strings.Split(username, "/")
	if len(s) != 2 {
		return "", "",
			fmt.Errorf("failed to get owner and environment name from the ssh username")
	}
	return s[0], s[1], nil
}
