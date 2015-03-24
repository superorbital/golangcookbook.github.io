package strings_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Strings Suite")
}
