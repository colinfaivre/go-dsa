package training

// 7.8 `M` Binary Tree Level Order Traversal

/*** @LEETCODE https://leetcode.com/problems/binary-tree-level-order-traversal/
Given the root of a binary tree, return the level order traversal of its nodes' values.
(i.e., from left to right, level by level).
---
Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
---
Example 2:
Input: root = [1]
Output: [[1]]
---
Example 3:
Input: root = []
Output: []
---
Constraints:
The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000
***/

/*** @SOLUTION https://www.youtube.com/watch?v=6ZnyEApgFYg
***/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil { return [][]int{} }
    var result [][]int
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        levelSize := len(queue)
        var level []int
        for i := 0; i < levelSize; i++ {
            // dequeue node to append it to level
            node := queue[0]
            queue = queue[1:]
            level = append(level, node.Val)
            
            if node.Left != nil { queue = append(queue, node.Left) } // enqueue node.Left
            if node.Right != nil { queue = append(queue, node.Right) } // enqueue node.Right
        }
        result = append(result, level)
    }

    return result
}
