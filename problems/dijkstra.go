/*
Dikstra algorithm
@DOCUMENTARY https://www.youtube.com/watch?v=mLEOZO1GwVc
@MEDIUM https://medium.com/@balajeraam/dijkstra-the-person-algorithm-5016ebe9468
*/

package problems

import (
	"github.com/colinfaivre/go-dsa/datastructures"
)

// O(mn): Computes a map of all shortest path starting from the given vertex
func DijkstraNaive(graph *datastructures.Graph, start_vertex int) map[int]int {
	shortest_paths := map[int]int{start_vertex: 0}

	for len(shortest_paths) != len(graph.GetVertices()) {

		border := [][4]int{}

		for v := range shortest_paths {

			for w := range graph.GetVertices()[v].Next {
				_, ok := shortest_paths[w]
				if !ok {
					border = append(border, [4]int{v, w, shortest_paths[v], graph.GetVertices()[v].Next[w]})
				}

			}

		}

		min := [4]int{0, 0, 0, 1000000}

		for _, item := range border {
			if item[2]+item[3] < min[2]+min[3] {
				min = item
			}
		}

		shortest_paths[min[1]] = shortest_paths[min[0]] + min[3]
	}

	return shortest_paths
}

// O(mlogn): Computes a map of all shortest path starting from the given vertex
func DijkstraHeap(graph *datastructures.Graph, start_vertex int) map[int]int {
	shortest_paths := map[int]int{}
	for k := range graph.GetVertices() {
		shortest_paths[k] = 1000000
	}
	heap := datastructures.NewHeap(true)

	heap.Insert(datastructures.Item{Id: start_vertex, Value: 0})
	shortest_paths[start_vertex] = 0

	for heap.Size() > 0 {
		v := heap.ExtractFrom(0).Id

		for w, weight := range graph.GetVertices()[v].Next {
			if shortest_paths[w] > shortest_paths[v]+weight {
				shortest_paths[w] = shortest_paths[v] + weight
				heap.Insert(datastructures.Item{Id: w, Value: shortest_paths[w]})
			}
		}
	}

	return shortest_paths
}
