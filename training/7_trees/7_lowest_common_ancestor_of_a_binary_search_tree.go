package training

// 7.7 `M` Lowest Common Ancestor of a Binary Search Tree

/*** @LEETCODE https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.
According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants
(where we allow a node to be a descendant of itself).”
---
Example 1:
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
---
Example 2:
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.
---
Example 3:
Input: root = [2,1], p = 2, q = 1
Output: 2
---
Constraints:
The number of nodes in the tree is in the range [2, 10^5].
-10^9 <= Node.val <= 10^9
All Node.val are unique.
p != q
p and q will exist in the BST.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=gs2LMfuOR9ks
***/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || !isSubtree(root, p) || !isSubtree(root, q) { return nil }

    lowestLeft := lowestCommonAncestor(root.Left, p, q)
    lowestRight := lowestCommonAncestor(root.Right, p, q)

    if lowestLeft != nil {
        return lowestLeft
    } else if lowestRight != nil {
        return lowestRight
    } else {
        return root
    }
}

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
