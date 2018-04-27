package ginserver_test

import (
	"net/http"

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

	Context("provided a valid gin engine and api path", func() {
		It("registers the necessary api routes", func() {
			expected := map[string]string{
				"/":    "GET",
				"/doc": "GET",
			}

			for _, route := range engine.Routes() {
				if _, ok := expected[route.Path]; !ok {
					continue
				}

				delete(expected, route.Path)
			}

			Expect(expected).To(HaveLen(0))
		})

		It("serves api doc from /doc", func() {
			resp := doRequest(engine, "GET", "/doc", nil, nil)
			Expect(resp).NotTo(BeNil())
			Expect(resp.Code).To(Equal(http.StatusMovedPermanently))
		})

	})
})
