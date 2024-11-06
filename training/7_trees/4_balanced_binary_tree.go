package training

// 7.4 `E` Balanced Binary Tree

/*** @LEETCODE https://leetcode.com/problems/balanced-binary-tree/
Given a binary tree, determine if it is height-balanced
---
Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: true
---
Example 2:
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
---
Example 3:
Input: root = []
Output: true
---
Constraints:
The number of nodes in the tree is in the range [0, 5000].
-10^4 <= Node.val <= 10^4
***/

/*** @SOLUTION https://www.youtube.com/watch?v=QfJsau0ItOY
***/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    res := isBalanced(root.Left) && isBalanced(root.Right)
    diff := depth(root.Left) - depth(root.Right)
    if math.Abs(diff) > 1 {
        return false
    } else {
        return true && res
    }
}

func depth(root *TreeNode) float64 {
    if root == nil {
        return 0
    }
    depthLeft := depth(root.Left)
    depthRight := depth(root.Right)
    return max(depthLeft, depthRight) + 1
}

