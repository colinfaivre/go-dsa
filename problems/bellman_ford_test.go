package problems_test

import (
	"github.com/colinfaivre/go-dsa/datastructures"
	"github.com/colinfaivre/go-dsa/parsing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BellmanFord()", func() {
	Context("Huge graph 1", func() {
		It("computes the assignement right clustering", func() {
			huge_graph := datastructures.NewGraph(true)
			adj_list, _ := parsing.ReadIntegers3TuplesFromFile("../test/data/g1")
			huge_graph.Add3TuplesEdges(adj_list)

			Expect(len(huge_graph.GetVertices())).To(Equal(1000))
		})
	})

	Context("Huge graph 2", func() {
		It("computes the assignement right clustering", func() {
			huge_graph := datastructures.NewGraph(true)
			adj_list, _ := parsing.ReadIntegers3TuplesFromFile("../test/data/g2")
			huge_graph.Add3TuplesEdges(adj_list)

			Expect(len(huge_graph.GetVertices())).To(Equal(1000))
		})
	})

	Context("Huge graph 3", func() {
		It("computes the assignement right clustering", func() {
			huge_graph := datastructures.NewGraph(true)
			adj_list, _ := parsing.ReadIntegers3TuplesFromFile("../test/data/g3")
			huge_graph.Add3TuplesEdges(adj_list)

			Expect(len(huge_graph.GetVertices())).To(Equal(1000))
		})
	})
})
