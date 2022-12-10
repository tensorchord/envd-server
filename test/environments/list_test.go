// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package environments

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"github.com/google/uuid"

	"github.com/tensorchord/envd-server/client"
	"github.com/tensorchord/envd-server/test/util"
)

var _ = Describe("environments", Ordered, func() {
	identityToken := uuid.New().String()
	logger := logrus.WithField("test-case", "environment-list").
		WithField("identity-token", identityToken)
	logger.Debug("Running test cases")
	// TODO(gaocegege): Add creation test case.
	Describe("with newly created environments", func() {
		logger.Debug(identityToken)
		srv, err := util.NewServer(util.NewPod("test", identityToken))
		Expect(err).Should(BeNil())
		cli, err := client.NewClientWithOpts(client.WithJWTToken(identityToken, ""))
		Expect(err).Should(BeNil())

		go func() {
			err := srv.Run()
			Expect(err).Should(BeNil())
		}()
		It("should get the newly created environments", func() {
			resp, err := cli.EnvironmentList(context.TODO())
			Expect(err).Should(BeNil())
			Expect(len(resp.Items)).Should(Equal(1))
		})
	})
})
