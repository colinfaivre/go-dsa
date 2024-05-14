package datastructures_test

import (
	"github.com/colinfaivre/go-dsa/datastructures"
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
		var adj_list = [][][2]int{{{2, 1}, {3, 4}}, {{3, 2}, {4, 6}}, {{4, 3}}}

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
	})

	// Context("Huge graph", func() {
	// 	huge_graph := datastructures.NewGraph(true)
	// 	edge_list, _ := parsing.ReadIntegersTuplesFromFile("../test/data/directed_graph")
	// 	huge_graph.AddEdges(edge_list)

	// 	It("should compute the last vertex with correct data", func() {
	// 		expected := &datastructures.Vertex{
	// 			Next: map[int]bool{
	// 				233422: true,
	// 				233423: true,
	// 				233424: true,
	// 				233425: true,
	// 				233426: true,
	// 				233427: true,
	// 				233428: true,
	// 				233429: true,
	// 				233430: true,
	// 				233431: true,
	// 				233432: true,
	// 				233433: true,
	// 				233434: true,
	// 				233435: true,
	// 				233436: true,
	// 			},
	// 		}
	// 		received := huge_graph.GetVertices()[875713]

	// 		Expect(received).To(Equal(expected))
	// 	})
	// })
})
