package training

// 9.12 `M` Redundant Connection

/*** @LEETCODE https://leetcode.com/problems/redundant-connection/
In this problem, a tree is an undirected graph that is connected and has no cycles.
You are given a graph that started as a tree with n nodes labeled from 1 to n, with one additional edge added.
The added edge has two different vertices chosen from 1 to n, and was not an edge that already existed.
The graph is represented as an array edges of length n where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the graph.
Return an edge that can be removed so that the resulting graph is a tree of n nodes.
If there are multiple answers, return the answer that occurs last in the input.
---
Example 1:
Input: edges = [[1,2],[1,3],[2,3]]
Output: [2,3]
---
Example 2:
Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
Output: [1,4]
---
Constraints:
n == edges.length
3 <= n <= 1000
edges[i].length == 2
1 <= ai < bi <= edges.length
ai != bi
There are no repeated edges.
The given graph is connected.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=FXWRE67PLL0
***/

// Find the root of a node with path compression
func find(parent []int, x int) int {
    if parent[x] != x {
        parent[x] = find(parent, parent[x]) // Path compression
    }
    return parent[x]
}

// Union two nodes and return false if they are already connected (cycle detected)
func union(parent, rank []int, x, y int) bool {
    rootX := find(parent, x)
    rootY := find(parent, y)

    if rootX == rootY {
        return false // Cycle detected
    }

    // Union by rank
    if rank[rootX] > rank[rootY] {
        parent[rootY] = rootX
    } else if rank[rootX] < rank[rootY] {
        parent[rootX] = rootY
    } else {
        parent[rootY] = rootX
        rank[rootX]++
    }

    return true
}

func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    parent := make([]int, n+1) // Parent array for union-find
    rank := make([]int, n+1)  // Rank array for union-find

    // Initialize the parent of each node to itself
    for i := 1; i <= n; i++ {
        parent[i] = i
    }

    for _, edge := range edges {
        x, y := edge[0], edge[1]
        if !union(parent, rank, x, y) {
            return edge // This edge forms a cycle
        }
    }

    return []int{} // Should never reach here for a valid input
}

