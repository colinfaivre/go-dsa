package training

// 6.2 `E` Merge Two Sorted Lists

/*** @LEETCODE https://leetcode.com/problems/merge-two-sorted-lists/
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
---
Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
---
Example 2:
Input: list1 = [], list2 = []
Output: []
---
Example 3:
Input: list1 = [], list2 = [0]
Output: [0]
---
Constraints:
The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=XIdigk956u0
O(n) solution:
- init a dummy ListNode ref
- init tail with dummy
- loop while list1 and list2 exist:
    - if list1.Val < list2.Val:    
        - set tail.Next to list1
        - set list1 to list1.Next
    - else:
        - set tail.Next to list2
        - set list2 to list2.Next
    - set tail to tail.Next
    - if list1 exists set tail.Next to list1
    - if list2 exists set tail.Next to list2
    - return dummy.Next
***/

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }
        tail = tail.Next
    }
    if list1 != nil {
        tail.Next = list1
    }
    if list2 != nil {
        tail.Next = list2
    }

    return dummy.Next
}
