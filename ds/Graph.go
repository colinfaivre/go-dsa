package ds

type Graph struct {
	is_directed bool
	vertices    map[int]Vertex
}

type Vertex struct {
	is_visited     bool
	finishing_time int
	leader         int
	edges          map[int]bool
}

func NewGraph(is_directed bool) *Graph {
	return &Graph{
		is_directed: is_directed,
		vertices:    map[int]Vertex{},
	}
}

func (graph *Graph) GetVertices() map[int]Vertex {
	return graph.vertices
}

func (graph *Graph) AddVertex(v int) {
	_, ok := graph.vertices[v]

	if !ok {
		graph.vertices[v] = Vertex{
			edges:          map[int]bool{},
			is_visited:     false,
			finishing_time: 0,
			leader:         0,
		}
	}
}

func (graph *Graph) AddEdge(v, w int) {
	graph.AddVertex(v)
	graph.AddVertex(w)

	graph.vertices[v].edges[w] = true
	if !graph.is_directed {
		graph.vertices[w].edges[v] = true
	}
}
