package ds

import (
	"testing"
)

func TestGraph(t *testing.T) {
	graph := NewGraph(true)
	graph.AddEdge(1, 2)
	graph.AddEdge(3, 4)
	graph.AddEdge(3, 5)
	graph.AddEdge(4, 6)

	t.Logf("log %v", graph)

	received := graph.GetVertices()

	if received[0] != 1 {
		t.Errorf("graph.GetVertices() expected [1,2,3,4,5,6] but got %v", received)
	}

	if received[5] != 6 {
		t.Errorf("graph.GetVertices() expected [1,2,3,4,5,6] but got %v", received)
	}
}
