// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package syncthing_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/tensorchord/envd-server/pkg/syncthing"
)

func TestSyncthing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Syncthing Suite")
}

var _ = Describe("Syncthing", func() {
	BeforeEach(func() {
	})

	Describe("Syncthing", func() {
		It("Generates configuration string", func() {
			configByte, err := (syncthing.GetConfigByte(syncthing.InitConfig()))
			configStr := string(configByte)
			Expect(err).To(BeNil())
			Expect(configStr).To(ContainSubstring("<configuration"))
			Expect(configStr).To(ContainSubstring("<apikey>envd</apikey>"))
			Expect(configStr).To(ContainSubstring("<address>0.0.0.0:8384</address>"))
			Expect(configStr).To(ContainSubstring("<globalAnnounceEnabled>false</globalAnnounceEnabled>"))
		})
	})
})
