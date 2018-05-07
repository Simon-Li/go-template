package interface_test

import (
	"testing"

	. "github.com/go-template/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInterface(t *testing.T) {
	RegisterFailHandler(Fail)
	TestPackage(t, "Interface Suite")
}
