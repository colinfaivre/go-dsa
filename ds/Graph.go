package ds

type Graph struct {
	is_directed bool
	vertices    map[int]*Vertex
}

type Vertex struct {
	is_explored    bool
	finishing_time int
	leader         int
	next           map[int]bool
	prev           map[int]bool
}

func NewGraph(is_directed bool) *Graph {
	return &Graph{
		is_directed: is_directed,
		vertices:    map[int]*Vertex{},
	}
}

func (graph *Graph) GetVertices() map[int]*Vertex {
	return graph.vertices
}

func (graph *Graph) AddVertex(v int) {
	_, ok := graph.vertices[v]

	if !ok {
		graph.vertices[v] = &Vertex{
			next:           map[int]bool{},
			prev:           map[int]bool{},
			is_explored:    false,
			finishing_time: 0,
			leader:         0,
		}
	}
}

func (graph *Graph) AddEdge(v, w int) {
	graph.AddVertex(v)
	graph.AddVertex(w)

	graph.vertices[v].next[w] = true
	if !graph.is_directed {
		graph.vertices[w].next[v] = true
	} else {
		graph.vertices[w].prev[v] = true
	}
}

func (graph *Graph) AddEdges(edge_list [][2]int) {
	for _, edge := range edge_list {
		graph.AddEdge(edge[0], edge[1])
	}
}

func (graph *Graph) IsEplored(v int) bool {
	return graph.vertices[v].is_explored
}

func (graph *Graph) AreExplored(vertices []int) bool {
	for _, v := range vertices {
		if !graph.vertices[v].is_explored {
			return false
		}
	}
	return true
}

func (graph *Graph) DFS(start_vertex int) {
	graph.vertices[start_vertex].is_explored = true

	for edge := range graph.vertices[start_vertex].next {
		if !graph.vertices[edge].is_explored {
			graph.DFS(edge)
		}
	}
}

func (graph *Graph) ReverseDFS(start_vertex int) {
	graph.vertices[start_vertex].is_explored = true

	for edge := range graph.vertices[start_vertex].prev {
		if !graph.vertices[edge].is_explored {
			graph.ReverseDFS(edge)
		}
	}
}

func (graph *Graph) BFS(start_vertex int) {
	graph.vertices[start_vertex].is_explored = true
	queue := Queue{}
	queue.Enqueue(start_vertex)

	for !queue.IsEmpty() {
		v := queue.Dequeue()
		for w := range graph.vertices[v].next {
			if !graph.vertices[w].is_explored {
				graph.vertices[w].is_explored = true
				queue.Enqueue(w)
			}
		}
	}
}
