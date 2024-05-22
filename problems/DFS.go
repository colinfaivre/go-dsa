package problems

import (
	"github.com/colinfaivre/go-dsa/datastructures"
)

func CallDFS(
	g *datastructures.Graph,
	start_vertex int,
) (map[int]bool, []int) {
	is_explored := map[int]bool{}
	search_path := []int{}

	DFS(g, start_vertex, is_explored, &search_path)

	return is_explored, search_path
}

func DFS(
	g *datastructures.Graph,
	start_vertex int,
	is_explored map[int]bool,
	search_path *[]int,
) {
	is_explored[start_vertex] = true
	*search_path = append(*search_path, start_vertex)

	for v := range g.GetVertices()[start_vertex].Next {
		if !is_explored[v] {
			DFS(g, v, is_explored, search_path)
		}
	}
}
