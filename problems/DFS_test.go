package problems_test

import (
	"github.com/colinfaivre/go-dsa/datastructures"
	"github.com/colinfaivre/go-dsa/problems"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DFS()", func() {
	var adj_list = [][][2]int{{{3, 4}, {2, 1}}, {{4, 6}, {3, 2}}, {{4, 3}}}

	It("running DFS() from vertex 1 should explore vertices 1, 2, 3, 4", func() {
		graph := datastructures.NewGraph(true)
		graph.AddEdges(adj_list)
		explored, search_path := problems.CallDFS(graph, 1)

		Expect(explored).To(Equal(map[int]bool{1: true, 2: true, 3: true, 4: true}))
		Expect(search_path[0]).To(Equal(1))
	})
})
