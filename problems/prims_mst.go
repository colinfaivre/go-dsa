package problems

import (
	"github.com/colinfaivre/go-dsa/datastructures"
)

const MaxInt = int(^uint(0) >> 1)

// Brute force algo
func PrimsMSTNaive(g *datastructures.Graph) int {
	mst_cost := 0
	visited_vertices := map[int]datastructures.Vertex{}
	// @TODO get random vertex instead of 1 ?
	visited_vertices[1] = *g.GetVertices()[1]

	for len(visited_vertices) != len(g.GetVertices()) {
		cheapest_edge := MaxInt
		cheapest_v_id := 0

		for u_id := range visited_vertices {
			for v_id := range visited_vertices[u_id].Next {
				_, ok := visited_vertices[v_id]
				if !ok && cheapest_edge > visited_vertices[u_id].Next[v_id] {
					cheapest_edge = visited_vertices[u_id].Next[v_id]
					cheapest_v_id = v_id
				}
			}
		}

		mst_cost += cheapest_edge
		visited_vertices[cheapest_v_id] = *g.GetVertices()[cheapest_v_id]
	}

	return mst_cost
}
