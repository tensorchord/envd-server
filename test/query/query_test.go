// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package query

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/tensorchord/envd-server/pkg/query"
	"github.com/tensorchord/envd-server/pkg/query/mock"
	"github.com/tensorchord/envd-server/pkg/service/user"
)

var _ = Describe("Mock test for db", func() {
	When("When change user data", func() {
		It("should work", func() {
			username := "test"
			pwd := "pwd"

			hashed, err := user.GenerateHashedSaltPassword([]byte(pwd))
			Expect(err).Should(BeNil())
			ctrl := gomock.NewController(GinkgoT())
			m := mock.NewMockQuerier(ctrl)
			m.EXPECT().CreateUser(
				context.Background(),
				gomock.All(),
			).Return(
				username, nil,
			)
			m.EXPECT().GetUser(
				context.Background(), username).Return(
				query.User{
					LoginName:    username,
					PasswordHash: string(hashed),
				}, nil,
			)
			userService := user.NewService(m, "secret", 0)
			_, err = userService.Register(context.Background(), username, pwd)
			Expect(err).NotTo(HaveOccurred())
			exists, _, err := userService.Login(context.Background(), username, pwd, true)
			Expect(exists).To(BeTrue())
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
