package ds

type Graph struct {
	is_directed bool
	vertices    map[int]bool
	adj_list    map[int][]int
}

func NewGraph(is_directed bool) *Graph {
	return &Graph{
		is_directed: is_directed,
		vertices:    map[int]bool{},
		adj_list:    map[int][]int{},
	}
}

func (graph *Graph) GetVertices() map[int]bool {
	return graph.vertices
}

func (graph *Graph) GetAdjList() map[int][]int {
	return graph.adj_list
}

func (graph *Graph) AddVertex(v int) {
	_, ok := graph.vertices[v]

	if !ok {
		graph.vertices[v] = true
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
