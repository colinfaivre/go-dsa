package problems_test

import (
	"github.com/colinfaivre/go-dsa/parsing"
	"github.com/colinfaivre/go-dsa/problems"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HuffmanCoding()", func() {
	It("returns a Huffman tree", func() {
		char_frequencies, _ := parsing.ReadIntegersFromFile("../test/data/huffman")
		huffman_tree := problems.HuffmanCoding(char_frequencies)

		Expect(huffman_tree).NotTo(BeNil())
	})
})
