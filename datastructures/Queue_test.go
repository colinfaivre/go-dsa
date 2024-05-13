package datastructures_test

import (
	"github.com/colinfaivre/go-dsa/datastructures"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Queue", func() {
	Context("Initially", func() {
		queue := datastructures.Queue{}

		It("Has an empty slice of integers", func() {
			Expect(queue.IsEmpty()).To(Equal(true))
		})
	})

	Context("When Enqueue() is called", func() {
		queue := datastructures.Queue{}
		queue.Enqueue(42)

		It("adds an item at the end of the queue", func() {
			Expect(queue.GetItems()[len(queue.GetItems())-1]).To(Equal(42))
		})
	})

	Context("When Dequeue() is called on a one element queue", func() {
		queue := datastructures.Queue{}
		queue.Enqueue(42)
		removed := queue.Dequeue()

		It("returns the dequeued integer", func() {
			Expect(removed).To(Equal(42))
		})

		It("should be empty", func() {
			Expect(queue.IsEmpty()).To(Equal(true))
		})
	})

	Context("Enqueue and Dequeue Scenario", func() {
		queue := datastructures.Queue{}
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)
		queue.Dequeue()
		queue.Dequeue()

		It("should have 1", func() {
			Expect(queue.GetItems()).To(Equal([]int{3}))
		})
	})
})
