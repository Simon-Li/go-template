package signal_test

import (
	"testing"

	. "github.com/go-template/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSignal(t *testing.T) {
	RegisterFailHandler(Fail)
	TestPackage(t, "Signal Suite")
}
