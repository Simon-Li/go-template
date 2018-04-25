package gin_server_test

import (
	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestGinServer", func() {
	var (
		engine *gin.Engine
	)

	It("hello world", func() {
		str := "hello world"

		Expect(str).To(ContainSubstring("hello"))
	})
})
