package problems

import (
	"github.com/colinfaivre/go-dsa/datastructures"
)

func CallBFS(
	g *datastructures.Graph,
	start_vertex int,
) (map[int]bool, []int) {
	is_explored := map[int]bool{}
	search_path := []int{}

	DFS(g, start_vertex, is_explored, &search_path)

	return is_explored, search_path
}

func BFS(
	g *datastructures.Graph,
	start_vertex int,
	is_explored map[int]bool,
	search_path *[]int,
) {
	is_explored[start_vertex] = true
	queue := datastructures.Queue{}
	queue.Enqueue(start_vertex)

	for !queue.IsEmpty() {
		v := queue.Dequeue()
		for w := range g.GetVertices()[v].Next {
			if !is_explored[w] {
				is_explored[w] = true
				queue.Enqueue(w)
				*search_path = append(*search_path, w)
			}
		}
	}
}
