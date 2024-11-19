package training

// 7.11 `M` Validate Binary Search Tree

/*** @LEETCODE https://leetcode.com/problems/validate-binary-search-tree/
Given the root of a binary tree, determine if it is a valid binary search tree (BST).
A valid BST is defined as follows:
- The left
- subtree
- of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.
---
Example 1:
Input: root = [2,1,3]
Output: true
---
Example 2:
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
---
Constraints:
The number of nodes in the tree is in the range [1, 10^4].
-2^31 <= Node.val <= 2^31 - 1
***/

/*** @SOLUTION https://www.youtube.com/watch?v=s6ATEkipzow
***/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return validBST(root, math.MinInt, math.MaxInt)
}

func validBST(root *TreeNode, left, right int) bool {
    if root == nil { return true }
    if root.Val >= right || root.Val <= left { return false }

    isValidLeft := validBST(root.Left, left, root.Val)
    isValidRight := validBST(root.Right, root.Val, right)
    return isValidLeft && isValidRight
}
