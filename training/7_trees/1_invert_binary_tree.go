package training

// 7.1 `E` Invert Binary Tree

/*** @LEETCODE https://leetcode.com/problems/invert-binary-tree/
Given the root of a binary tree, invert the tree, and return its root.
---
Example 1:
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
---
Example 2:
Input: root = [2,1,3]
Output: [2,3,1]
---
Example 3:
Input: root = []
Output: []
---
Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
***/

/*** @SOLUTION https://www.youtube.com/watch?v=OnSn2XEQ4MY
***/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    temp := root.Left
    root.Left = invertTree(root.Right)
    root.Right = invertTree(temp)
    return root
}
