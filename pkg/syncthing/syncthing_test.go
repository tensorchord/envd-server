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
			configStr, err := syncthing.GetConfigString(syncthing.InitConfig())
			Expect(err).To(BeNil())
			Expect(configStr).To(ContainSubstring("<configuration"))
			Expect(configStr).To(ContainSubstring("<apikey>envd</apikey>"))
			Expect(configStr).To(ContainSubstring("<address>0.0.0.0:8384</address>"))
			Expect(configStr).To(ContainSubstring("<globalAnnounceEnabled>false</globalAnnounceEnabled>"))
		})
	})
})
