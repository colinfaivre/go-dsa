package datastructures_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/colinfaivre/go-dsa/datastructures"
)

var _ = Describe("LinkedList", func() {
	Context("AddFirst()", func() {
		It("should print the LinkedList in correct order", func() {
			linked_list := datastructures.LinkedList{}
			node1 := &datastructures.Node{Data: 1}
			node2 := &datastructures.Node{Data: 2}
			node3 := &datastructures.Node{Data: 3}
			node4 := &datastructures.Node{Data: 4}
			node5 := &datastructures.Node{Data: 5}

			linked_list.AddFirst(node1)
			linked_list.AddFirst(node2)
			linked_list.AddFirst(node3)
			linked_list.AddFirst(node4)
			linked_list.AddFirst(node5)

			Expect(linked_list.ToString()).To(Equal("54321"))
		})
	})
})
