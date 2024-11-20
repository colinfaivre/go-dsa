package training

// 7.12 `M` Kth Smallest Element In a Bst

/*** @LEETCODE https://leetcode.com/problems/kth-smallest-element-in-a-bst/
Given the root of a binary search tree, and an integer k,
return the kth smallest value (1-indexed) of all the valuesof the nodes in the tree.
---
Example 1:
Input: root = [3,1,4,null,2], k = 1
Output: 1
---
Example 2:
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
---
Constraints:
The number of nodes in the tree is n.
1 <= k <= n <= 10^4
0 <= Node.val <= 10^4
***/

/*** @SOLUTION https://www.youtube.com/watch?v=5LUXSvjmGCw
***/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    var values []int
    inOrderTraverse(root, &values)
    return values[k-1]
}

func inOrderTraverse(root *TreeNode, values *[]int) {
    if root == nil { return }
    inOrderTraverse(root.Left, values)
    *values = append(*values, root.Val)
    inOrderTraverse(root.Right, values)
}
