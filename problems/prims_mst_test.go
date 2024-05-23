package problems_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PrimsMST", func() {
	Context("Huge graph", func() {
		//huge_graph := datastructures.NewGraph(false)
		// @TODO write a parser to get adjlist
		//adj_list, _ := parsing.GetUndirectedWeightedGraphData("../test/data/undirected_weighted_graph")
		//huge_graph.AddEdges(adj_list)

		It("PrimsMST()", func() {
			Expect(true).To(Equal(true))
		})
	})
})
