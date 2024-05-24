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

func (g *Graph) GetVertices() map[int]*Vertex {
	return g.vertices
}

func (g *Graph) addVertex(v int) {
	_, ok := g.vertices[v]

	if !ok {
		g.vertices[v] = &Vertex{
			Next: map[int]int{},
		}
	}
}

func (g *Graph) AddEdge(v, w, weight int) {
	g.addVertex(v)
	g.addVertex(w)

	g.vertices[v].Next[w] = weight
	if !g.is_directed {
		g.vertices[w].Next[v] = weight
	}
}

func (g *Graph) AddEdges(adj_list [][][2]int) {
	for i, edge_list := range adj_list {
		for _, edge := range edge_list {
			g.AddEdge(i+1, edge[0], edge[1])
		}
	}
}

func (g *Graph) Add3TuplesEdges(adj_list [][3]int) {
	for _, e := range adj_list {
		g.AddEdge(e[0], e[1], e[2])
	}
}
