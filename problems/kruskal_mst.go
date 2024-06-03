package problems

import (
	"sort"

	"github.com/colinfaivre/go-dsa/datastructures"
)

// @MEDIUM https://yuminlee2.medium.com/kruskals-algorithm-minimum-spanning-tree-db96e91d0aed

// Computes the minimum spanning tree cost of given undirected edges
func KruskalsMST(edges [][3]int, numOfNodes int) (int, [][]int) {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	uf := datastructures.NewUnionFind(numOfNodes)

	minWeight := 0
	edgeCount := 0
	MSTedges := [][]int{}
	for _, edge := range edges {
		node1, node2, weight := edge[0], edge[1], edge[2]
		if uf.Union(node1, node2) {
			minWeight += weight
			edgeCount += 1
			MSTedges = append(MSTedges, []int{node1, node2, weight})

			if edgeCount == numOfNodes-1 {
				return minWeight, MSTedges
			}
		}
	}
	return minWeight, MSTedges
}
