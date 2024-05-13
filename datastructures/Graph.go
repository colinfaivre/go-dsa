package datastructures

import "github.com/colinfaivre/go-dsa/algorithms"

type Graph struct {
	is_directed           bool
	vertices              map[int]*Vertex
	search_path           []int
	curr_time             int
	curr_leader           int
	finish_time_to_vertex map[int]int
	is_explored           map[int]bool
	scc_sizes             map[int]int
}

type Vertex struct {
	Leader int
	Next   map[int]bool
	Prev   map[int]bool
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
		scc_sizes:             map[int]int{},
	}
}

func (graph *Graph) RestetSearchPath() {
	graph.search_path = []int{}
}

func (graph *Graph) GetSearchPath() []int {
	return graph.search_path
}

func (graph *Graph) GetSCCSizes() map[int]int {
	return graph.scc_sizes
}

func (graph *Graph) GetVertices() map[int]*Vertex {
	return graph.vertices
}

func (graph *Graph) AddVertex(v int) {
	_, ok := graph.vertices[v]

	if !ok {
		graph.vertices[v] = &Vertex{
			Next:   map[int]bool{},
			Prev:   map[int]bool{},
			Leader: 0,
		}
	}
}

func (graph *Graph) AddEdge(v, w int) {
	graph.AddVertex(v)
	graph.AddVertex(w)

	graph.vertices[v].Next[w] = true
	if !graph.is_directed {
		graph.vertices[w].Next[v] = true
	} else {
		graph.vertices[w].Prev[v] = true
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
	graph.vertices[start_vertex].Leader = graph.curr_leader
	graph.scc_sizes[graph.curr_leader]++
	graph.search_path = append(graph.search_path, start_vertex)

	for v := range graph.vertices[start_vertex].Next {
		if !graph.is_explored[v] {
			graph.DFS(v)
		}
	}
}

func (graph *Graph) ReverseDFS(start_vertex int) {
	graph.is_explored[start_vertex] = true
	graph.search_path = append(graph.search_path, start_vertex)

	for v := range graph.vertices[start_vertex].Prev {
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
		for w := range graph.vertices[v].Next {
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

func (graph *Graph) GetTheFiveBiggestSCC() []int {
	scc_size_list := make([]int, 0, len(graph.scc_sizes))

	for _, val := range graph.scc_sizes {
		scc_size_list = append(scc_size_list, val)
	}

	scc_size_list = algorithms.MergeSort(scc_size_list)
	return scc_size_list[len(scc_size_list)-5:]
}
