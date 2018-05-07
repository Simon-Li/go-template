package interface_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test interface handling", func() {

	Context("test basic interface handling works", func() {
		It("handles basic interface assertion", func() {
			var a interface{}
			a = (uint32)(1)
			b := a.(uint32)
			Expect(b).To(Equal(a))
		})
	})
})
