package ds

import "slices"

type Graph struct {
	is_directed bool
	vertices    []int
	adj_list    map[int][]int
}

func NewGraph(is_directed bool) *Graph {
	return &Graph{
		is_directed: is_directed,
		vertices:    []int{},
		adj_list:    map[int][]int{},
	}
}

func (graph *Graph) GetVertices() []int {
	return graph.vertices
}

func (graph *Graph) GetAdjList() map[int][]int {
	return graph.adj_list
}

func (graph *Graph) AddVertex(v int) {
	if !slices.Contains(graph.vertices, v) {
		graph.vertices = append(graph.vertices, v)
		graph.adj_list[v] = []int{}
	}
}

func (graph *Graph) AddEdge(v, w int) {
	if graph.adj_list[v] == nil {
		graph.AddVertex(v)
	}
	if graph.adj_list[w] == nil {
		graph.AddVertex(w)
	}
	graph.adj_list[v] = append(graph.adj_list[v], w)
	if !graph.is_directed {
		graph.adj_list[w] = append(graph.adj_list[w], v)
	}
}
