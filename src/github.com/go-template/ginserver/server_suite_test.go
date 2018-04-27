package ginserver_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-template/ginserver"
	. "github.com/go-template/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinEngine(t *testing.T) {
	RegisterFailHandler(Fail)
	TestPackage(t, "Gin Engine Suite")
}

func NewTestGinEngine() *gin.Engine {

	return ginserver.Engine()
}

func doRequest(engine http.Handler, method, path string, headers map[string]string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		panic(err)
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	resp := httptest.NewRecorder()
	engine.ServeHTTP(resp, req)
	return resp
}
