package ds

import (
	"testing"

	"github.com/colinfaivre/go-dsa/utils"
)

func TestGraph(t *testing.T) {
	// Small dataset
	graph := NewGraph(true)
	graph.AddEdge(1, 2)
	graph.AddEdge(3, 4)
	graph.AddEdge(3, 5)
	graph.AddEdge(4, 6)

	expected_vertices := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true}
	received_vertices := graph.GetVertices()

	expected_adj_list := map[int][]int{1: {2}, 2: {}, 3: {4, 5}, 4: {6}, 5: {}, 6: {}}
	received_adj_list := graph.GetAdjList()

	if expected_vertices[1] != received_vertices[1] {
		t.Errorf("graph.GetVertices() expected %v but received %v", expected_vertices, received_vertices)
	}

	if expected_vertices[5] != received_vertices[5] {
		t.Errorf("graph.GetVertices() expected %v but received %v", expected_vertices, received_vertices)
	}

	if expected_adj_list[1][0] != 2 && expected_adj_list[6] != nil {
		t.Errorf("graph.GetVertices() expected %v but received %v", expected_adj_list, received_adj_list)
	}

	// Huge dataset
	huge_graph := NewGraph(true)
	integer_tupples, _ := utils.ReadIntegersTuplesFromFile("../data/scc")
	n := len(integer_tupples)

	for i := 0; i < n; i++ {
		huge_graph.AddEdge(integer_tupples[i][0], integer_tupples[i][1])
	}

	expected_huge_adj_list := []int{233422, 233423, 233424, 233425, 233426, 233427, 233428, 233429, 233430, 233431, 233432, 233433, 233434, 233435, 233436}
	received_huge_adj_list := huge_graph.GetAdjList()[875713]

	if expected_huge_adj_list[0] != received_huge_adj_list[0] {
		t.Errorf("huge_graph.GetAdjList() expected %v but received %v", expected_adj_list, received_adj_list)
	}
}
