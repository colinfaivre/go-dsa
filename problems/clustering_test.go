package problems_test

import (
	"github.com/colinfaivre/go-dsa/datastructures"
	"github.com/colinfaivre/go-dsa/parsing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Clustering", func() {
	It("computes the assignement right clustering", func() {
		huge_graph := datastructures.NewGraph(false)
		adj_list, _ := parsing.ReadIntegers3TuplesFromFile("../test/data/clustering")
		huge_graph.Add3TuplesEdges(adj_list)

		Expect(len(huge_graph.GetVertices())).To(Equal(500))
	})
})
