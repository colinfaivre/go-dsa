package datastructures

// @WIKI https://en.wikipedia.org/wiki/Graph_(abstract_data_type)

type Graph struct {
	is_directed bool
	vertices    map[int]*Vertex
	search_path []int
}

type Vertex struct {
	Next map[int]int
}

func NewGraph(is_directed bool) *Graph {
	return &Graph{
		is_directed: is_directed,
		vertices:    map[int]*Vertex{},
		search_path: []int{},
	}
}

func (graph *Graph) GetVertices() map[int]*Vertex {
	return graph.vertices
}

func (graph *Graph) addVertex(v int) {
	_, ok := graph.vertices[v]

	if !ok {
		graph.vertices[v] = &Vertex{
			Next: map[int]int{},
		}
	}
}

func (graph *Graph) AddEdge(v, w, weight int) {
	graph.addVertex(v)
	graph.addVertex(w)

	graph.vertices[v].Next[w] = weight
	if !graph.is_directed {
		graph.vertices[w].Next[v] = weight
	}
}

func (graph *Graph) AddEdges(adj_list [][][2]int) {
	for i, edge_list := range adj_list {
		for _, edge := range edge_list {
			graph.AddEdge(i+1, edge[0], edge[1])
		}
	}
}
