package ds

import (
	"testing"

	"github.com/colinfaivre/go-dsa/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGraph(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graph Suite")
}

var _ = Describe("Graph", func() {
	Context("Initially directed", func() {
		graph := NewGraph(true)

		It("is directed", func() {
			Expect(graph.is_directed).Should(Equal(true))
		})

		It("has an empty map of vertices", func() {
			Expect(graph.GetVertices()).Should(Equal(map[int]*Vertex{}))
		})
	})

	Context("When adding an edge", func() {
		graph := NewGraph(true)
		graph.AddEdge(1, 2)

		It("should add two vertices", func() {
			Expect(1).Should(BeKeyOf(graph.vertices))
			Expect(2).Should(BeKeyOf(graph.vertices))
		})

		It("shoud add an edge to the first vertex", func() {
			Expect(2).Should(BeKeyOf(graph.vertices[1].next))
		})
	})

	Context("Small graph", func() {
		var edge_list = [][2]int{{1, 4}, {2, 8}, {3, 6}, {4, 7}, {5, 2}, {6, 9}, {7, 1}, {8, 6}, {8, 5}, {9, 3}, {9, 7}}

		It("should correctly add multiple vertices", func() {
			graph := NewGraph(true)
			graph.AddEdges(edge_list)

			expected := map[int]*Vertex{
				1: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					next:           map[int]bool{4: true},
					prev:           map[int]bool{7: true},
				},
				2: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					next:           map[int]bool{8: true},
					prev:           map[int]bool{5: true},
				},
				3: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					next:           map[int]bool{6: true},
					prev:           map[int]bool{9: true},
				},
				4: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					next:           map[int]bool{7: true},
					prev:           map[int]bool{1: true},
				},
				5: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					next:           map[int]bool{2: true},
					prev:           map[int]bool{8: true},
				},
				6: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					next:           map[int]bool{9: true},
					prev:           map[int]bool{3: true, 8: true},
				},
				7: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					next:           map[int]bool{1: true},
					prev:           map[int]bool{4: true, 9: true},
				},
				8: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					next:           map[int]bool{6: true, 5: true},
					prev:           map[int]bool{2: true},
				},
				9: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					next:           map[int]bool{3: true, 7: true},
					prev:           map[int]bool{6: true},
				},
			}
			received := graph.GetVertices()

			Expect(expected).Should(Equal(received))
		})

		It("running DFS() from vertex 1 should explore vertices 1, 4, 7", func() {
			graph := NewGraph(true)
			graph.AddEdges(edge_list)
			graph.DFS(1)

			Expect(graph.AreExplored([]int{4, 7})).To(BeTrue())
			Expect(graph.AreExplored([]int{3, 5})).To(BeFalse())
			Expect(graph.search_path).To(Equal([]int{1, 4, 7}))
		})

		It("running ReverseDFS() from vertex 6 should explore vertices 6, 3, 8, 2, 5, 9", func() {
			graph := NewGraph(true)
			graph.AddEdges(edge_list)
			graph.ReverseDFS(6)

			Expect(graph.AreExplored([]int{8, 2, 5, 3, 9})).To(BeTrue())
			Expect(graph.AreExplored([]int{1, 4, 7})).To(BeFalse())
			Expect(graph.search_path).To(Equal([]int{6, 3, 8, 2, 5, 9}))
		})

		It("running DFS() from vertex 6 should explore vertices 6, 9, 3, 7, 1, 4", func() {
			graph := NewGraph(true)
			graph.AddEdges(edge_list)
			graph.DFS(6)

			Expect(graph.AreExplored([]int{3, 9, 1, 4, 7})).To(BeTrue())
			Expect(graph.AreExplored([]int{8, 2, 5})).To(BeFalse())
			Expect(graph.search_path).To(Equal([]int{6, 9, 3, 7, 1, 4}))
		})
	})

	Context("Huge graph", func() {
		huge_graph := NewGraph(true)
		edge_list, _ := utils.ReadIntegersTuplesFromFile("../data/scc")
		huge_graph.AddEdges(edge_list)

		It("should compute the last vertex with correct data", func() {
			expected := &Vertex{
				is_explored:    false,
				finishing_time: 0,
				leader:         0,
				next: map[int]bool{
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
				prev: map[int]bool{},
			}
			received := huge_graph.GetVertices()[875713]

			Expect(received).Should(Equal(expected))
		})

		It("can run DFS", func() {
			huge_graph.DFS(2)
		})

		It("can run ReverseDFS", func() {
			huge_graph.ReverseDFS(2)
		})
	})
})
