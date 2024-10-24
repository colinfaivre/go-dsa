package training

// 6.6 `M` Add Two Numbers

/*** @LEETCODE https://leetcode.com/problems/add-two-numbers/
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
---
Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
---
Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]
---
Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
---
Constraints:
The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=wgFPrzTjm7s
***/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy
    mem := 0

    for l1 != nil || l2 != nil {
        curr.Next = &ListNode{}
        curr = curr.Next

        var sum int
        if l1 == nil {
            sum = l2.Val + mem
        } else if l2 == nil {
            sum = l1.Val + mem
        } else {
            sum = l1.Val + l2.Val + mem
        }

        if sum < 10 {
            curr.Val = sum
            mem = 0
        } else {
            curr.Val = sum - 10
            mem = 1
        }

        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }
    }

    if mem == 1 {
        curr.Next = &ListNode{Val: 1}
    }

    return dummy.Next
}
