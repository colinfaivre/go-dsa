package training

// 6.3 `M` Reorder List

/*** @LEETCODE https://leetcode.com/problems/reorder-list/
You are given the head of a singly linked-list. The list can be represented as:
L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.
---
Example 1:
Input: head = [1,2,3,4]
Output: [1,4,2,3]
---
Example 2:
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
---
Constraints:
The number of nodes in the list is in the range [1, 5 * 104].
1 <= Node.val <= 1000
***/

/*** @SOLUTION https://www.youtube.com/watch?v=S5bfdUTrKLM
***/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReorderList(head *ListNode)  {
    // find middle
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // reverse second half
    second := slow.Next
    var prev *ListNode = nil
    slow.Next = nil
    for second != nil {
        tmp := second.Next
        second.Next = prev
        prev = second
        second = tmp
    }
    // merge 2 halfs
    first := head
    second = prev
    for second != nil {
        tmp1, tmp2 := first.Next, second.Next
        first.Next = second
        second.Next = tmp1
        first, second = tmp1, tmp2
    }
}
