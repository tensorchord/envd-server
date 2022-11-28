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
	"github.com/tensorchord/envd-server/pkg/service"
)

var _ = Describe("Mock test for db", func() {
	When("When change user data", func() {
		It("should work", func() {
			ctrl := gomock.NewController(GinkgoT())
			m := mock.NewMockQuerier(ctrl)
			m.EXPECT().CreateUser(context.Background(), query.CreateUserParams{IdentityToken: "test", PublicKey: []byte("whoami")}).Return(query.User{}, nil)
			m.EXPECT().GetUser(context.Background(), "test").Return(query.User{IdentityToken: "test", PublicKey: []byte("whoami")}, nil)
			userService := service.NewUserService(m)
			err := userService.Register("test", []byte("whoami"))
			Expect(err).NotTo(HaveOccurred())
			exists, err := userService.Auth("test")
			Expect(exists).To(BeTrue())
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
