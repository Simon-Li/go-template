package main_test

import (
	"testing"

	. "github.com/go-template/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	TestPackage(t, "Main Suite")
}
