package training

// 7.6 `E` Subtree of Another Tree

/*** @LEETCODE https://leetcode.com/problems/subtree-of-another-tree/
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root
with the same structure and node values of subRoot and false otherwise.
A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants.
The tree tree could also be considered as a subtree of itself.
---
Example 1:
Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true
---
Example 2:
Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false
---
Constraints:
The number of nodes in the root tree is in the range [1, 2000].
The number of nodes in the subRoot tree is in the range [1, 1000].
-10^4 <= root.val <= 10^4
-10^4 <= subRoot.val <= 10^4
***/

/*** @SOLUTION https://www.youtube.com/watch?v=E36O5SWp-LE
***/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil { return false }

    isSame := isSameTree(root, subRoot)
    isSameLeft := isSameTree(root.Left, subRoot)
    isSameRight := isSameTree(root.Right, subRoot)
    isSubLeft := isSubtree(root.Left, subRoot)
    isSubRight := isSubtree(root.Right, subRoot)
    return isSameLeft || isSameRight || isSubLeft || isSubRight || isSame
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil { return true }
    if p == nil || q == nil { return false }
    if p.Val != q.Val { return false }

    isSameLeft := isSameTree(p.Left, q.Left)
    isSameRight := isSameTree(p.Right, q.Right)
    return isSameLeft && isSameRight
}
