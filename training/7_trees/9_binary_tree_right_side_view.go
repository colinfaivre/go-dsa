package training

// 7.9 `M` Binary Tree Right Side View

/*** @LEETCODE https://leetcode.com/problems/binary-tree-right-side-view/
Given the root of a binary tree, imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.
---
Example 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
---
Example 2:
Input: root = [1,null,3]
Output: [1,3]
---
Example 3:
Input: root = []
Output: []
---
Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
***/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil { return []int{} }
    var res []int
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        levelSize := len(queue)
        var level []int

        for i := 0; i < levelSize; i++ {
            // dequeue node
            node := queue[0]
            queue = queue[1:]
            level = append(level, node.Val)
            if node.Left != nil {
                // enqueue node.Left
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                // enqueue node.Right
                queue = append(queue, node.Right)
            }
        }

        res = append(res, level[len(level) - 1])
    }

    return res
}

/*** @SOLUTION https://www.youtube.com/watch?v=d4zLyf32e3I
***/
