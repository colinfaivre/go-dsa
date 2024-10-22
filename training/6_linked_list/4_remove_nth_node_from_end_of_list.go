package training

// 6.4 `M` Remove Nth Node From End of List

/*** @LEETCODE https://leetcode.com/problems/remove-nth-node-from-end-of-list/
Given the head of a linked list, remove the nth node from the end of the list and return its head.
---
Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
---
Example 2:
Input: head = [1], n = 1
Output: []
---
Example 3:
Input: head = [1,2], n = 1
Output: [1]
---
Constraints:
The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
***/

/*** @SOLUTION https://www.youtube.com/watch?v=XVuQxVej6y8
***/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    listLength := 0
    countPointer := head
    for countPointer != nil {
        listLength++
        countPointer = countPointer.Next
    }
    toRemove := listLength - n

    if toRemove > 0 {
        pointer := head
        for i := 0; i < toRemove - 1; i++ {
            pointer = pointer.Next
        }
        pointer.Next = pointer.Next.Next
    } else if listLength == 1 && n == 1{
        head = nil
    } else if listLength == n {
        head = head.Next
    }

    return head
}
