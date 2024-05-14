package datastructures

// WIKI https://en.wikipedia.org/wiki/Graph_(abstract_data_type)

type Graph struct {
	is_directed    bool
	vertices       map[int]*Vertex
	search_path    []int
	is_explored    map[int]bool
	shortest_paths map[int]int
}

type Vertex struct {
	Next map[int]int
}

func NewGraph(is_directed bool) *Graph {
	return &Graph{
		is_directed:    is_directed,
		vertices:       map[int]*Vertex{},
		search_path:    []int{},
		is_explored:    map[int]bool{},
		shortest_paths: map[int]int{},
	}
}

func (graph *Graph) GetShortestPaths() map[int]int {
	return graph.shortest_paths
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
			Next: map[int]int{},
		}
	}
}

func (graph *Graph) AddEdge(v, w, weight int) {
	graph.AddVertex(v)
	graph.AddVertex(w)

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

// https://medium.com/@balajeraam/dijkstra-the-person-algorithm-5016ebe9468
func (graph *Graph) Dijkstra(start_vertex int) {
	graph.shortest_paths[start_vertex] = 0

	for len(graph.shortest_paths) != len(graph.vertices) {

		border := [][3]int{}

		for v := range graph.shortest_paths {

			for w := range graph.vertices[v].Next {
				_, ok := graph.shortest_paths[w]
				if !ok {
					border = append(border, [3]int{v, w, graph.vertices[v].Next[w]})
				}

			}

		}

		min := [3]int{0, 0, 1000000}

		for _, item := range border {
			if item[2] < min[2] {
				min = item
			}
		}

		graph.shortest_paths[min[1]] = graph.shortest_paths[min[0]] + min[2]
	}
}
