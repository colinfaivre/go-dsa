package ds

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stack Suite")
}

var _ = Describe("Stack", func() {
	Context("Initially", func() {
		stack := Stack{}

		It("Has an empty slice of integers", func() {
			Expect(stack.IsEmpty()).Should(Equal(true))
		})
	})

	Context("When Push() is called", func() {
		stack := Stack{}
		stack.Push(42)

		It("adds an item at the end of the stack", func() {
			Expect(stack.items[len(stack.items)-1]).Should(Equal(42))
		})
	})

	Context("When Pop() is called on a one element stack", func() {
		stack := Stack{}
		stack.Push(42)
		removed := stack.Pop()

		It("returns the popped integer", func() {
			Expect(removed).Should(Equal(42))
		})

		It("should be empty", func() {
			Expect(stack.IsEmpty()).Should(Equal(true))
		})
	})
})
