// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package query

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

func TestEnvdServer(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	RegisterFailHandler(Fail)
	RunSpecs(t, "query mock Suite")
}
