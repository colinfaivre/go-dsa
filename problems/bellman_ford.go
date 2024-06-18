package problems

import "github.com/colinfaivre/go-dsa/datastructures"

// @HELPFUL https://reintech.io/blog/bellman-ford-algorithm-in-go
// @HELPFUL https://github.com/pprcht/shortestpaths/blob/master/go/bellman-ford.go

// O(mn): Computes a map of all shortest paths starting from the given vertex
// works with positive and negative edge costs
// finds negative cycles
func BellManFord(g *datastructures.Graph, start_vertex int) ([]int, bool) {
	res := []int{}
	hasNegativeCycle := false

	return res, hasNegativeCycle
}
