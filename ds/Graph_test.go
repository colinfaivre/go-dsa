package ds

import (
	"testing"

	"github.com/colinfaivre/go-dsa/utils"
)

func TestGraph(t *testing.T) {
	// graph := NewGraph(true)
	// graph.AddEdge(1, 2)
	// graph.AddEdge(3, 4)
	// graph.AddEdge(3, 5)
	// graph.AddEdge(4, 6)

	// t.Logf("log %v", graph)

	// received := graph.GetVertices()

	// if received[0] != 1 {
	// 	t.Errorf("graph.GetVertices() expected [1,2,3,4,5,6] but got %v", received)
	// }

	// if received[5] != 6 {
	// 	t.Errorf("graph.GetVertices() expected [1,2,3,4,5,6] but got %v", received)
	// }

	huge_graph := NewGraph(true)
	integer_tupples, _ := utils.ReadIntegersTuplesFromFile("../data/scc")
	n := len(integer_tupples)

	for i := 0; i < n; i++ {
		huge_graph.AddEdge(integer_tupples[i][0], integer_tupples[i][1])
	}

	t.Logf("log huge_graph - %v", huge_graph.adj_list[875714])
}
