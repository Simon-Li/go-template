package ginserver_test

import (
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
