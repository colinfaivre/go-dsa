package problems_test

import (
	"github.com/colinfaivre/go-dsa/datastructures"
	"github.com/colinfaivre/go-dsa/parsing"
	"github.com/colinfaivre/go-dsa/problems"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PrimsMST", func() {
	Context("Huge graph", func() {
		huge_graph := datastructures.NewGraph(false)
		adj_list, _ := parsing.ReadIntegers3TuplesFromFile("../test/data/undirected_weighted_graph")
		huge_graph.Add3TuplesEdges(adj_list)

		It("PrimsMST()", func() {
			Expect(len(huge_graph.GetVertices())).To(Equal(500))
		})

		It("Computes the right MST cost", func() {
			Expect(problems.PrimsMSTNaive(huge_graph)).To(Equal(-3612829))
		})
	})
})
