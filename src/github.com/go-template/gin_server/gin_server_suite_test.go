package gin_server_test

import (
	"testing"

	. "github.com/go-template/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinEngine(t *testing.T) {
	RegisterFailHandler(Fail)
	TestPackage(t, "Gin Engine Suite")
}
