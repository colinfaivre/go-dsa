package datastructures

// WIKI https://en.wikipedia.org/wiki/Graph_(abstract_data_type)

type Graph struct {
	is_directed bool
	vertices    map[int]*Vertex
	search_path []int
	is_explored map[int]bool
}

type Vertex struct {
	Next map[int]bool
}

func NewGraph(is_directed bool) *Graph {
	return &Graph{
		is_directed: is_directed,
		vertices:    map[int]*Vertex{},
		search_path: []int{},
		is_explored: map[int]bool{},
	}
}

func (graph *Graph) RestetSearchPath() {
	graph.search_path = []int{}
}

func (graph *Graph) GetSearchPath() []int {
	return graph.search_path
}

func (graph *Graph) GetVertices() map[int]*Vertex {
	return graph.vertices
}

func (graph *Graph) AddVertex(v int) {
	_, ok := graph.vertices[v]

	if !ok {
		graph.vertices[v] = &Vertex{
			Next: map[int]bool{},
		}
	}
}

func (graph *Graph) AddEdge(v, w int) {
	graph.AddVertex(v)
	graph.AddVertex(w)

	graph.vertices[v].Next[w] = true
	if !graph.is_directed {
		graph.vertices[w].Next[v] = true
	}
}

func (graph *Graph) AddEdges(edge_list [][2]int) {
	for _, edge := range edge_list {
		graph.AddEdge(edge[0], edge[1])
	}
}

func (graph *Graph) AreExplored(vertices []int) bool {
	for _, v := range vertices {
		if !graph.is_explored[v] {
			return false
		}
	}
	return true
}

func (graph *Graph) DFS(start_vertex int) {
	graph.is_explored[start_vertex] = true
	graph.search_path = append(graph.search_path, start_vertex)

	for v := range graph.vertices[start_vertex].Next {
		if !graph.is_explored[v] {
			graph.DFS(v)
		}
	}
}

func (graph *Graph) BFS(start_vertex int) {
	graph.is_explored[start_vertex] = true
	queue := Queue{}
	queue.Enqueue(start_vertex)

	for !queue.IsEmpty() {
		v := queue.Dequeue()
		for w := range graph.vertices[v].Next {
			if !graph.is_explored[w] {
				graph.is_explored[w] = true
				queue.Enqueue(w)
				graph.search_path = append(graph.search_path, w)
			}
		}
	}
}
