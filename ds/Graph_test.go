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
			Expect(2).Should(BeKeyOf(graph.vertices[1].edges))
		})
	})

	Context("Small graph", func() {
		// Small dataset
		graph := NewGraph(true)
		graph.AddEdge(1, 4)
		graph.AddEdge(2, 8)
		graph.AddEdge(3, 6)
		graph.AddEdge(4, 7)
		graph.AddEdge(5, 2)
		graph.AddEdge(6, 9)
		graph.AddEdge(7, 1)
		graph.AddEdge(8, 6)
		graph.AddEdge(8, 5)
		graph.AddEdge(9, 3)
		graph.AddEdge(9, 7)

		It("should correctly add multiple vertices", func() {
			expected := map[int]*Vertex{
				1: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					edges:          map[int]bool{4: true},
				},
				2: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					edges:          map[int]bool{8: true},
				},
				3: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					edges:          map[int]bool{6: true},
				},
				4: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					edges:          map[int]bool{7: true},
				},
				5: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					edges:          map[int]bool{2: true},
				},
				6: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					edges:          map[int]bool{9: true},
				},
				7: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					edges:          map[int]bool{1: true},
				},
				8: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					edges:          map[int]bool{6: true, 5: true},
				},
				9: {
					is_explored:    false,
					finishing_time: 0,
					leader:         0,
					edges:          map[int]bool{3: true, 7: true},
				},
			}
			received := graph.GetVertices()

			Expect(expected).Should(Equal(received))
		})

		It("running DFS() from vertex 1 should explore vertices 4 and 7 ", func() {
			graph.DFS(1)

			Expect(graph.vertices[4].is_explored).Should(BeTrue())
			Expect(graph.vertices[7].is_explored).Should(BeTrue())
			Expect(graph.vertices[3].is_explored).Should(BeFalse())
			Expect(graph.vertices[5].is_explored).Should(BeFalse())
		})
	})

	Context("Huge graph", func() {
		huge_graph := NewGraph(true)
		integer_tupples, _ := utils.ReadIntegersTuplesFromFile("../data/scc")
		n := len(integer_tupples)
		for i := 0; i < n; i++ {
			huge_graph.AddEdge(integer_tupples[i][0], integer_tupples[i][1])
		}

		It("should compute the last vertex with correct data", func() {
			expected := &Vertex{
				is_explored:    false,
				finishing_time: 0,
				leader:         0,
				edges: map[int]bool{
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

			Expect(received).Should(Equal(expected))
		})
	})
})
