package ds

type Graph struct {
	is_directed           bool
	vertices              map[int]*Vertex
	search_path           []int
	curr_time             int
	curr_leader           int
	finish_time_to_vertex map[int]int
	is_explored           map[int]bool
}

type Vertex struct {
	leader int
	next   map[int]bool
	prev   map[int]bool
}

func NewGraph(is_directed bool) *Graph {
	return &Graph{
		is_directed:           is_directed,
		vertices:              map[int]*Vertex{},
		search_path:           []int{},
		curr_time:             0,
		curr_leader:           0,
		finish_time_to_vertex: map[int]int{},
		is_explored:           map[int]bool{},
	}
}

func (graph *Graph) RestetSearchPath() {
	graph.search_path = []int{}
}

func (graph *Graph) GetVertices() map[int]*Vertex {
	return graph.vertices
}

func (graph *Graph) AddVertex(v int) {
	_, ok := graph.vertices[v]

	if !ok {
		graph.vertices[v] = &Vertex{
			next:   map[int]bool{},
			prev:   map[int]bool{},
			leader: 0,
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

func (graph *Graph) IsExplored(v int) bool {
	return graph.is_explored[v]
}

func (graph *Graph) AreExplored(vertices []int) bool {
	for _, v := range vertices {
		if !graph.is_explored[v] {
			return false
		}
	}
	return true
}

func (graph *Graph) InitSearchOrder() {
	graph.search_path = []int{}
}

func (graph *Graph) DFS(start_vertex int) {
	graph.is_explored[start_vertex] = true
	graph.vertices[start_vertex].leader = graph.curr_leader
	graph.search_path = append(graph.search_path, start_vertex)

	for v := range graph.vertices[start_vertex].next {
		if !graph.is_explored[v] {
			graph.DFS(v)
		}
	}
}

func (graph *Graph) ReverseDFS(start_vertex int) {
	graph.is_explored[start_vertex] = true
	graph.search_path = append(graph.search_path, start_vertex)

	for v := range graph.vertices[start_vertex].prev {
		if !graph.is_explored[v] {
			graph.ReverseDFS(v)
		}
	}

	graph.curr_time++
	graph.finish_time_to_vertex[graph.curr_time] = start_vertex
}

func (graph *Graph) BFS(start_vertex int) {
	graph.is_explored[start_vertex] = true
	queue := Queue{}
	queue.Enqueue(start_vertex)

	for !queue.IsEmpty() {
		v := queue.Dequeue()
		for w := range graph.vertices[v].next {
			if !graph.is_explored[w] {
				graph.is_explored[w] = true
				queue.Enqueue(w)
				graph.search_path = append(graph.search_path, w)
			}
		}
	}
}

func (graph *Graph) Kosaraju() {

	for i := len(graph.vertices); i >= 1; i-- {
		_, ok := graph.vertices[i]
		if ok && !graph.is_explored[i] {
			graph.ReverseDFS(i)
		}
	}

	graph.is_explored = map[int]bool{}

	for i := len(graph.finish_time_to_vertex); i >= 1; i-- {
		_, ok := graph.vertices[graph.finish_time_to_vertex[i]]
		if ok && !graph.is_explored[graph.finish_time_to_vertex[i]] {
			graph.curr_leader = i
			graph.DFS(graph.finish_time_to_vertex[i])
		}
	}
}
