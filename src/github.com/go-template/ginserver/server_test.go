package ginserver_test

import (
	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestGinServer", func() {
	var (
		engine *gin.Engine
	)

	BeforeEach(func() {
		engine = NewTestGinEngine()
	})

	It("hello world", func() {
		str := "hello world"

		Expect(str).To(ContainSubstring("hello"))
	})
})
