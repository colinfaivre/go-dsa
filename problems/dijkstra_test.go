package problems_test

import (
	"github.com/colinfaivre/go-dsa/datastructures"
	"github.com/colinfaivre/go-dsa/parsing"
	"github.com/colinfaivre/go-dsa/problems"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dijkstra", func() {
	Context("Small graph", func() {
		It("Dijkstra(1)", func() {
			adj_list := [][][2]int{{{3, 4}, {2, 1}}, {{4, 6}, {3, 2}}, {{4, 3}}}
			graph := datastructures.NewGraph(true)
			graph.AddEdges(adj_list)
			shortest_path := problems.DijkstraNaive(graph, 1)

			Expect(shortest_path[4]).To(Equal(6))
		})

		It("Dijkstra(1) on course test case with big shortcut", func() {
			adj_list := [][][2]int{{{2, 1}, {8, 2}}, {{1, 1}, {3, 1}}, {{2, 1}, {4, 1}}, {{3, 1}, {5, 1}}, {{4, 1}, {6, 1}}, {{5, 1}, {7, 1}}, {{6, 1}, {8, 1}}, {{7, 1}, {1, 2}}}
			graph := datastructures.NewGraph(true)
			graph.AddEdges(adj_list)
			shortest_path := problems.DijkstraNaive(graph, 1)

			Expect(shortest_path[1]).To(Equal(0))
			Expect(shortest_path[2]).To(Equal(1))
			Expect(shortest_path[3]).To(Equal(2))
			Expect(shortest_path[4]).To(Equal(3))
			Expect(shortest_path[5]).To(Equal(4))
			Expect(shortest_path[6]).To(Equal(4))
			Expect(shortest_path[7]).To(Equal(3))
			Expect(shortest_path[8]).To(Equal(2))
		})
	})

	Context("Huge graph", func() {
		huge_graph := datastructures.NewGraph(true)
		adj_list, _ := parsing.GetWeightedGraphData("../test/data/directed_weighted_graph")
		huge_graph.AddEdges(adj_list)

		It("Dijkstra(1)", func() {
			sortest_path := problems.DijkstraNaive(huge_graph, 1)

			Expect(sortest_path[7]).To(Equal(2599))
			Expect(sortest_path[37]).To(Equal(2610))
			Expect(sortest_path[59]).To(Equal(2947))
			Expect(sortest_path[82]).To(Equal(2052))
			Expect(sortest_path[99]).To(Equal(2367))
			Expect(sortest_path[115]).To(Equal(2399))
			Expect(sortest_path[133]).To(Equal(2029))
			Expect(sortest_path[165]).To(Equal(2442))
			Expect(sortest_path[188]).To(Equal(2505))
			Expect(sortest_path[197]).To(Equal(3068))
		})
	})
})
