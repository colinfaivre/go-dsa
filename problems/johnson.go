package problems

import "github.com/colinfaivre/go-dsa/datastructures"

// O(mnlogn): Computes a map of all shortest paths for each (u, v) in Graph
// works with positive and negative edge costs
// finds negative cycles
func Johnson(g *datastructures.Graph) (map[int]int, bool) {
	res := map[int]int{}
	hasNegativeCycle := false

	// #1 O(n)		form G' by adding a new vertex s and a new edge (s, v)
	// 				with weight 0 for each v of G

	// #2 O(mn)		run BellMan-Ford on G' from s
	// 				if B-F detects a negative cost cycle in G', halt and report

	// #3 O(m)		for each v of G, define:
	//				Pv := length of a shortest (s, v) path in G'
	//				for each edge e = (u, v) of G, define:
	//				Ce' := Ce + Pu - Pv

	// #4 O(nmlogn)	for each u of G:
	// 				run Dijkstra in G with edge lengths Ce', from vertex u
	// 				to compute the shortest-path distance d'(u, v) for each v of G

	// #5 O(n^2)	for each pair (u, v) of G
	//				return the shortest-path distance:
	//				d(u, v) := d'(u, v) - Pu + Pv

	return res, hasNegativeCycle
}
