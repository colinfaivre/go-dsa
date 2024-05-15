package problems_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProblems(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Problems Suite")
}
