package datastructures_test

import (
	"github.com/colinfaivre/go-dsa/datastructures"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BST", func() {
	Context("Initially", func() {
		tree := datastructures.BST{}

		It("Has an empty root", func() {
			Expect(tree.GetRoot()).To(BeNil())
		})
	})

	Context("Insert()", func() {
		tree := datastructures.BST{}

		It("adds keys on right sides", func() {
			tree.Insert(2)
			tree.Insert(1)
			tree.Insert(3)

			Expect(tree.GetRoot().Key).To(Equal(2))
			Expect(tree.GetRoot().Left.Key).To(Equal(1))
			Expect(tree.GetRoot().Right.Key).To(Equal(3))
		})

		// @TODO handle if key is already there
	})

	Context("Search()", func() {
		tree := datastructures.BST{}

		It("finds a key if it is in the tree", func() {
			tree.Insert(2)
			tree.Insert(1)
			tree.Insert(3)

			Expect(tree.Search(3)).To(Equal(true))
			Expect(tree.Search(10)).To(Equal(false))
		})
	})

	Context("Min() / Max()", func() {
		tree := datastructures.BST{}

		It("returns the min and max values of the tree", func() {
			tree.Insert(2)
			tree.Insert(1)
			tree.Insert(3)

			Expect(tree.Min()).To(Equal(1))
			Expect(tree.Max()).To(Equal(3))
		})
	})
})
