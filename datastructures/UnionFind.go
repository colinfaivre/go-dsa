package datastructures

// @MEDIUM https://yuminlee2.medium.com/kruskals-algorithm-minimum-spanning-tree-db96e91d0aed

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(numOfNodes int) *UnionFind {
	// makeSet
	parent := make([]int, numOfNodes)
	size := make([]int, numOfNodes)
	for i := 0; i < numOfNodes; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{
		parent: parent,
		size:   size,
	}
}

func (uf *UnionFind) Find(node int) int {
	for node != uf.parent[node] {
		uf.parent[node] = uf.parent[uf.parent[node]]
		node = uf.parent[node]
	}
	return node
}

func (uf *UnionFind) Union(node1, node2 int) bool {
	root1 := uf.Find(node1)
	root2 := uf.Find(node2)

	if root1 == root2 {
		return false
	}

	if uf.size[root1] > uf.size[root2] {
		uf.parent[root2] = root1
		uf.size[root1] += 1
	} else {
		uf.parent[root1] = root2
		uf.size[root2] += 1
	}
	return true
}
