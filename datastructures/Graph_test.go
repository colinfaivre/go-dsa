package datastructures_test

import (
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
		graph.AddEdge(1, 2)

		It("should add two vertices", func() {
			Expect(1).To(BeKeyOf(graph.GetVertices()))
			Expect(2).To(BeKeyOf(graph.GetVertices()))
		})

		It("shoud add an edge to the first vertex", func() {
			Expect(2).To(BeKeyOf(graph.GetVertices()[1].Next))
		})
	})

	Context("Small graph", func() {
		var edge_list = [][2]int{{1, 4}, {2, 8}, {3, 6}, {4, 7}, {5, 2}, {6, 9}, {7, 1}, {8, 6}, {8, 5}, {9, 3}, {9, 7}}

		It("should correctly add multiple vertices", func() {
			graph := datastructures.NewGraph(true)
			graph.AddEdges(edge_list)

			expected := map[int]*datastructures.Vertex{
				1: {
					Next: map[int]bool{4: true},
				},
				2: {
					Next: map[int]bool{8: true},
				},
				3: {
					Next: map[int]bool{6: true},
				},
				4: {
					Next: map[int]bool{7: true},
				},
				5: {
					Next: map[int]bool{2: true},
				},
				6: {
					Next: map[int]bool{9: true},
				},
				7: {
					Next: map[int]bool{1: true},
				},
				8: {
					Next: map[int]bool{6: true, 5: true},
				},
				9: {
					Next: map[int]bool{3: true, 7: true},
				},
			}
			received := graph.GetVertices()

			Expect(expected).To(Equal(received))
		})

		It("running DFS() from vertex 1 and 2 should explore vertices 1, 4, 7, 2, 8, 6, 9, 3, 5", func() {
			graph := datastructures.NewGraph(true)
			graph.AddEdges(edge_list)
			graph.DFS(1)
			graph.DFS(2)

			Expect(graph.AreExplored([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})).To(BeTrue())
			Expect(graph.GetSearchPath()[0]).To(Equal(1))
		})
	})

	Context("Huge graph", func() {
		huge_graph := datastructures.NewGraph(true)
		edge_list, _ := parsing.ReadIntegersTuplesFromFile("../test/data/directed_graph")
		huge_graph.AddEdges(edge_list)

		It("should compute the last vertex with correct data", func() {
			expected := &datastructures.Vertex{
				Next: map[int]bool{
					233422: true,
					233423: true,
					233424: true,
					233425: true,
					233426: true,
					233427: true,
					233428: true,
					233429: true,
					233430: true,
					233431: true,
					233432: true,
					233433: true,
					233434: true,
					233435: true,
					233436: true,
				},
			}
			received := huge_graph.GetVertices()[875713]

			Expect(received).To(Equal(expected))
		})
	})
})
