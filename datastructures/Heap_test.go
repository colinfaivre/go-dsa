package datastructures_test

import (
	"github.com/colinfaivre/go-dsa/datastructures"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MaxHeap", func() {
	Context("Initially", func() {
		heap := datastructures.MaxHeap{}

		It("Has an empty slice of integers", func() {
			Expect(len(heap.GetArray())).To(Equal(0))
		})
	})

	Context("Insert()", func() {
		heap := datastructures.MaxHeap{}
		initialHeapLength := len(heap.GetArray())
		heap.Insert(42)

		It("adds an item to the heap", func() {
			Expect(len(heap.GetArray())).To(Equal(initialHeapLength + 1))
			Expect(heap.GetArray()[0]).To(Equal(42))
		})
	})

	Context("Heapify()", func() {
		heap := datastructures.MaxHeap{}
		heap.Heapify([]int{10, 20, 30})

		It("has 3 values", func() {
			Expect(len(heap.GetArray())).To(Equal(3))
		})

		It("has the max value at the root of the heap", func() {
			Expect(heap.GetArray()[0]).To(Equal(30))
		})
	})

	Context("Extract()", func() {
		heap := datastructures.MaxHeap{}
		heap.Heapify([]int{10, 20, 30, 40, 50, 60, 70, 80})
		heap.Extract(0)

		It("should have the biggest value as first element", func() {
			Expect(heap.GetArray()[0]).To(Equal(70))
		})
	})
})
