package problems

import "github.com/colinfaivre/go-dsa/datastructures"

// O(mn): Computes a map of all shortest paths starting from the given vertex
// works with positive and negative edge costs
// finds negative cycles
func BellManFord(g *datastructures.Graph, start_vertex int) (map[int]int, bool) {
	res := map[int]int{}
	hasNegativeCycle := false

	return res, hasNegativeCycle
}
