package datastructures_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/colinfaivre/go-dsa/datastructures"
)

var _ = Describe("LinkedList", func() {
	Context("AddAllFront()", func() {
		It("should add all items to the front of the linked list", func() {
			linked_list := datastructures.LinkedList{}
			linked_list.AddAllFront([]int{1, 2, 3, 4, 5})

			Expect(linked_list.ToString()).To(Equal("12345"))
		})
	})

	Context("AddFirst()", func() {
		It("should add one item at the front of the linked list", func() {
			linked_list := datastructures.LinkedList{}
			linked_list.AddAllFront([]int{1, 2, 3, 4, 5})
			original_size := linked_list.Size()
			linked_list.AddFirst(0)

			Expect(linked_list.PeekFirst()).To(Equal(0))
			Expect(linked_list.ToString()).To(Equal("012345"))
			Expect(linked_list.Size()).To(Equal(original_size + 1))
		})
	})

	Context("RemoveFirst()", func() {
		It("should remove and return the first item of the linked list", func() {
			linked_list := datastructures.LinkedList{}
			linked_list.AddAllFront([]int{1, 2, 3, 4, 5})
			original_size := linked_list.Size()
			linked_list.RemoveFirst()

			Expect(linked_list.PeekFirst()).To(Equal(2))
			Expect(linked_list.ToString()).To(Equal("2345"))
			Expect(linked_list.Size()).To(Equal(original_size - 1))
		})
	})

	Context("AddLast()", func() {
		It("should add one item at the back of the linked list", func() {
			linked_list := datastructures.LinkedList{}
			linked_list.AddAllFront([]int{1, 2, 3, 4, 5})
			original_size := linked_list.Size()
			linked_list.AddLast(6)
			linked_list.AddLast(7)

			Expect(linked_list.PeekLast()).To(Equal(7))
			Expect(linked_list.ToString()).To(Equal("1234567"))
			Expect(linked_list.Size()).To(Equal(original_size + 2))
		})
	})
})
