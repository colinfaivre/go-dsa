package training

// 6.1 `E` Reverse Linked List

/*** @LEETCODE https://leetcode.com/problems/reverse-linked-list/
Given the head of a singly linked list, reverse the list, and return the reversed list.
---
Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
---
Example 2:
Input: head = [1,2]
Output: [2,1]
---
Example 3:
Input: head = []
Output: []
---
Constraints:
The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000
***/

/*** @SOLUTION https://www.youtube.com/watch?v=G0_I-ZF0S38
Iterative T:O(n) S:O(1) solution:
- declare res as a pointer to a ListNode
- set pointer at head
- loop while pointer exists:
    - set temp to pointer.Next
    - set pointer.Next to res
    - set res to pointer
    - set pointer to temp
- return prev
***/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseList(head *ListNode) *ListNode {
    var res *ListNode
    pointer := head

    for pointer != nil {
        temp := pointer.Next
        pointer.Next = res
        res = pointer
        pointer = temp
    }

    return res
}
