package datastructures_test

import (
	"github.com/colinfaivre/go-dsa/datastructures"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Heap", func() {
	Context("MaxHeap Initially", func() {
		heap := datastructures.NewHeap(false)

		It("Has an empty slice of integers", func() {
			Expect(heap.Size()).To(Equal(0))
		})
	})

	Context("MaxHeap Insert()", func() {
		heap := datastructures.NewHeap(false)
		initialHeapLength := heap.Size()
		heap.Insert(datastructures.Item{Value: 42})

		It("adds an item to the heap", func() {
			Expect(heap.Size()).To(Equal(initialHeapLength + 1))
			Expect(heap.Peek()).To(Equal(datastructures.Item{Value: 42}))
		})
	})

	Context("MaxHeap Heapify()", func() {
		heap := datastructures.NewHeap(false)
		heap.Heapify([]int{
			10,
			20,
			30,
		})

		It("has 3 values", func() {
			Expect(heap.Size()).To(Equal(3))
		})

		It("has the max value at the root of the heap", func() {
			Expect(heap.Peek()).To(Equal(datastructures.Item{Value: 30}))
		})
	})

	Context("MaxHeap Extract()", func() {
		heap := datastructures.NewHeap(false)
		heap.Heapify([]int{
			10,
			20,
			30,
			40,
			50,
			60,
			70,
			80,
		})
		heap.ExtractFrom(0)

		It("should have the biggest value as first element", func() {
			Expect(heap.Peek()).To(Equal(datastructures.Item{Value: 70}))
		})
	})

	Context("MinHeap Heapify()", func() {
		heap := datastructures.NewHeap(true)
		heap.Heapify([]int{
			10,
			20,
			30,
		})

		It("has 3 values", func() {
			Expect(heap.Size()).To(Equal(3))
		})

		It("has the max value at the root of the heap", func() {
			Expect(heap.Peek()).To(Equal(datastructures.Item{Value: 10}))
		})
	})

	Context("MinHeap Extract()", func() {
		heap := datastructures.NewHeap(true)
		heap.Heapify([]int{
			10,
			20,
			30,
			40,
			50,
			60,
			70,
			80,
		})

		It("should have the smallest value as first element", func() {
			Expect(heap.ExtractFrom(0)).To(Equal(datastructures.Item{Value: 10}))
		})
	})
})
