package training

// 7.3 `E` Diameter of Binary Tree

/*** @LEETCODE https://leetcode.com/problems/diameter-of-binary-tree/
Given the root of a binary tree, return the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
This path may or may not pass through the root.
The length of a path between two nodes is represented by the number of edges between them.
---
Example 1:
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
---
Example 2:
Input: root = [1,2]
Output: 1
---
Constraints:
The number of nodes in the tree is in the range [1, 10^4].
-100 <= Node.val <= 100
***/

/*** @SOLUTION https://www.youtube.com/watch?v=bkxqA8Rfv04
***/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    currDiam := depth(root.Left) + depth(root.Right)
    leftDiam := diameterOfBinaryTree(root.Left)
    rightDiam := diameterOfBinaryTree(root.Right)
    return max(currDiam, leftDiam, rightDiam)
}

func depth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    depthLeft := depth(root.Left)
    depthRight := depth(root.Right)

    return 1 + max(depthLeft, depthRight)
}
