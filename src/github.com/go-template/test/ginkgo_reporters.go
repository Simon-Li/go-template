package test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"testing"
)

// Create junit style reporter for unit tests (required for Insights reporting)
func TestPackage(t *testing.T, suite string) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, suite, []Reporter{junitReporter})
}
