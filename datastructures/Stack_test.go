package datastructures_test

import (
	"github.com/colinfaivre/go-dsa/datastructures"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack", func() {
	Context("Initially", func() {
		stack := datastructures.Stack{}

		It("Has an empty slice of integers", func() {
			Expect(stack.IsEmpty()).To(Equal(true))
		})
	})

	Context("When Push() is called", func() {
		stack := datastructures.Stack{}
		stack.Push(42)

		It("adds an item at the end of the stack", func() {
			Expect(stack.GetItems()[len(stack.GetItems())-1]).To(Equal(42))
		})
	})

	Context("When Pop() is called on a one element stack", func() {
		stack := datastructures.Stack{}
		stack.Push(42)
		removed := stack.Pop()

		It("returns the popped integer", func() {
			Expect(removed).To(Equal(42))
		})

		It("should be empty", func() {
			Expect(stack.IsEmpty()).To(Equal(true))
		})
	})

	Context("Push and Pop Scenario", func() {
		stack := datastructures.Stack{}
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		stack.Pop()
		stack.Pop()

		It("should have 1", func() {
			Expect(stack.GetItems()).To(Equal([]int{1}))
		})
	})
})
