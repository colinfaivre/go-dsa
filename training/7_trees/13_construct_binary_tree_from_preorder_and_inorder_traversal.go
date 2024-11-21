package training

// 7.13 `M` Construct Binary Tree From Preorder And Inorder Traversal

/*** @LEETCODE https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree
and inorder is the inorder traversal of the same tree, construct and return the binary tree.
---
Example 1:
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
---
Example 2:
Input: preorder = [-1], inorder = [-1]
Output: [-1]
---
Constraints:
1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=ihj4IQGZ2zc
***/

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    root := &TreeNode{Val: preorder[0]}

    mid := 0
    for i, val := range inorder {
        if val == preorder[0] {
            mid = i
            break
        }
    }

    root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
    root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
    
    return root
}
