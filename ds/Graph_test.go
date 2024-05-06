package ds

import (
	"testing"

	"github.com/colinfaivre/go-dsa/utils"
)

func TestGraph(t *testing.T) {
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

	expected_vertices := map[int][]int{1: {4}, 2: {8}, 3: {6}, 4: {7}, 5: {2}, 6: {9}, 7: {1}, 8: {6, 5}, 9: {3, 7}}
	received_vertices := graph.GetVertices()
	graph.DFS(1)

	t.Logf("graph %v", graph.vertices[2])

	if !received_vertices[1].edges[4] && !received_vertices[9].edges[7] {
		t.Errorf("graph.GetVertices() expected %v but received %v", expected_vertices, received_vertices)
	}

	// Huge dataset
	huge_graph := NewGraph(true)
	integer_tupples, _ := utils.ReadIntegersTuplesFromFile("../data/scc")
	n := len(integer_tupples)

	for i := 0; i < n; i++ {
		huge_graph.AddEdge(integer_tupples[i][0], integer_tupples[i][1])
	}

	// t.Logf("huge_graph %v", huge_graph.vertices[875713])
	expected_huge_vertices := Vertex{is_explored: false, finishing_time: 0, leader: 0, edges: map[int]bool{233422: true, 233423: true, 233424: true, 233425: true, 233426: true, 233427: true, 233428: true, 233429: true, 233430: true, 233431: true, 233432: true, 233433: true, 233434: true, 233435: true, 233436: true}}
	received_huge_vertices := huge_graph.GetVertices()[875713]

	if !received_huge_vertices.edges[233422] || !received_huge_vertices.edges[233436] {
		t.Errorf("huge_graph.GetVertices() expected %v but received %v", expected_huge_vertices, received_huge_vertices)
	}
}
