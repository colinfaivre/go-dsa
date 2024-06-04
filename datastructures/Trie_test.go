package datastructures_test

import (
	"github.com/colinfaivre/go-dsa/datastructures"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trie", func() {
	Context("Initially", func() {
		trie := datastructures.NewTrie()

		It("Has a root containing a node size 26 array of children nodes", func() {
			Expect(len(trie.GetRoot().Children)).To(Equal(26))
			Expect(trie.GetRoot().Children[0]).To(BeNil())
			Expect(trie.GetRoot().Children[25]).To(BeNil())
		})

		It("Has a root containing a node with isEnd set to false", func() {
			Expect(trie.GetRoot().IsEnd).To(Equal(false))
		})
	})

	Context("Insert()", func() {
		trie := datastructures.NewTrie()
		trie.Insert("cat")

		It("takes a word and adds it into the Trie", func() {
			Expect(trie.GetRoot().Children[2].Children[0].Children[19]).NotTo((BeNil()))
		})
	})

	Context("Search()", func() {
		trie := datastructures.NewTrie()
		trie.Insert("cat")

		It("returns true if the searched word is in the Trie", func() {
			Expect(trie.Search("cat")).To(BeTrue())
		})

		It("returns false if the searched word is in the Trie", func() {
			Expect(trie.Search("ca")).To(BeFalse())
			Expect(trie.Search("dog")).To(BeFalse())
			Expect(trie.Search("cam")).To(BeFalse())
			Expect(trie.Search("catapult")).To(BeFalse())
		})
	})
})
