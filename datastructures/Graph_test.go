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
	})
})
