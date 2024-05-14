package datastructures_test

import (
	"fmt"

	"github.com/colinfaivre/go-dsa/datastructures"
	"github.com/colinfaivre/go-dsa/parsing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Graph", func() {
	Context("Initially", func() {
		graph := datastructures.NewGraph(true)

		It("has an empty map of vertices", func() {
			Expect(graph.GetVertices()).To(Equal(map[int]*datastructures.Vertex{}))
		})
	})

	Context("When adding an edge", func() {
		graph := datastructures.NewGraph(true)
		graph.AddEdge(1, 2, 50)

		It("should add two vertices", func() {
			Expect(1).To(BeKeyOf(graph.GetVertices()))
			Expect(2).To(BeKeyOf(graph.GetVertices()))
		})

		It("shoud add an edge to the first vertex with correct weight", func() {
			Expect(graph.GetVertices()[1].Next[2]).To(Equal(50))
		})
	})

	Context("When adding multiple edges", func() {
		graph := datastructures.NewGraph(true)
		graph.AddEdges([][][2]int{{{2, 50}, {3, 75}, {4, 100}}, {{1, 150}}})

		It("should add four vertices", func() {
			Expect(1).To(BeKeyOf(graph.GetVertices()))
			Expect(2).To(BeKeyOf(graph.GetVertices()))
			Expect(3).To(BeKeyOf(graph.GetVertices()))
			Expect(4).To(BeKeyOf(graph.GetVertices()))
		})

		It("shoud add an edges to the first vertex with correct weights", func() {
			Expect(graph.GetVertices()[1].Next[2]).To(Equal(50))
			Expect(graph.GetVertices()[1].Next[3]).To(Equal(75))
			Expect(graph.GetVertices()[1].Next[4]).To(Equal(100))
			Expect(graph.GetVertices()[2].Next[1]).To(Equal(150))
		})
	})

	Context("Small graph", func() {
		var adj_list = [][][2]int{{{3, 4}, {2, 1}}, {{4, 6}, {3, 2}}, {{4, 3}}}

		It("should correctly add multiple vertices", func() {
			graph := datastructures.NewGraph(true)
			graph.AddEdges(adj_list)

			expected := map[int]*datastructures.Vertex{
				1: {
					Next: map[int]int{2: 1, 3: 4},
				},
				2: {
					Next: map[int]int{3: 2, 4: 6},
				},
				3: {
					Next: map[int]int{4: 3},
				},
				4: {
					Next: map[int]int{},
				},
			}
			received := graph.GetVertices()

			Expect(expected).To(Equal(received))
		})

		It("running DFS() from vertex 1 should explore vertices 1, 2, 3, 4", func() {
			graph := datastructures.NewGraph(true)
			graph.AddEdges(adj_list)
			graph.DFS(1)

			Expect(graph.AreExplored([]int{1, 2, 3, 4})).To(BeTrue())
			Expect(graph.GetSearchPath()[0]).To(Equal(1))
		})

		It("Dijkstra(1)", func() {
			graph := datastructures.NewGraph(true)
			graph.AddEdges(adj_list)
			graph.Dijkstra(1)
			fmt.Println("result small", graph.GetShortestPaths())
			Expect(graph.GetShortestPaths()[4]).To(Equal(6))
		})

		It("Dijkstra(1) with course test case", func() {
			var adj_list = [][][2]int{{{2, 1}, {8, 2}}, {{1, 1}, {3, 1}}, {{2, 1}, {4, 1}}, {{3, 1}, {5, 1}}, {{4, 1}, {6, 1}}, {{5, 1}, {7, 1}}, {{6, 1}, {8, 1}}, {{7, 1}, {1, 2}}}
			graph := datastructures.NewGraph(true)
			graph.AddEdges(adj_list)
			graph.Dijkstra(1)
			Expect(graph.GetShortestPaths()[1]).To(Equal(0))
			Expect(graph.GetShortestPaths()[2]).To(Equal(1))
			Expect(graph.GetShortestPaths()[3]).To(Equal(2))
			Expect(graph.GetShortestPaths()[4]).To(Equal(3))
			Expect(graph.GetShortestPaths()[5]).To(Equal(4))
			// next ones failing because of a shortcut by 8
			Expect(graph.GetShortestPaths()[6]).To(Equal(4))
			Expect(graph.GetShortestPaths()[7]).To(Equal(3))
			Expect(graph.GetShortestPaths()[8]).To(Equal(2))
		})
	})

	Context("Huge graph", func() {
		huge_graph := datastructures.NewGraph(true)
		adj_list, _ := parsing.GetWeightedGraphData("../test/data/directed_weighted_graph")
		huge_graph.AddEdges(adj_list)
		// fmt.Println("huge graph", huge_graph)

		It("Dijkstra(1)", func() {
			huge_graph.Dijkstra(1)
			sp := huge_graph.GetShortestPaths()

			fmt.Println("result huge", huge_graph.GetShortestPaths())

			// 7,37,59,82,99,115,133,165,188,197
			fmt.Println("result huge", sp[7], sp[37], sp[59], sp[82], sp[99], sp[115], sp[133], sp[165], sp[188], sp[197])
			// 1st response 4685 2610 6222 2052 6893 2834 2029 4399 2633 4483
		})
	})
})
