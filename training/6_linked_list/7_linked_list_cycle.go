package training

// 6.7 `E` Linked List Cycle

/*** @LEETCODE https://leetcode.com/problems/linked-list-cycle/
Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to.
Note that pos is not passed as a parameter.
Return true if there is a cycle in the linked list. Otherwise, return false.
---
Example 1:
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
---
Example 2:
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
---
Example 3:
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
---
Constraints:
The number of the nodes in the list is in the range [0, 10^4].
-10^5 <= Node.val <= 10^5
pos is -1 or a valid index in the linked-list.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=gBTe7lFR3vc
O(n) solution - HashSet:
- init hashSet to store visited ListNode pointers
- loop while head exists:
  - if hashSet[head] exists return true
  - else add head to hashSet
  - update head to head.Next
- return false

O(1) solution - Two pointers:
- init slow and fast to head
- loop while fast and fast.Next exist:
  - set slow to slow.Next
  - set fast to fast.Next.Next
  if slow equals fast return true
- return false
***/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func HasCycle(head *ListNode) bool {
    hashSet := map[*ListNode]bool{}
    for head != nil {
        if hashSet[head] {
            return true
        } else {
            hashSet[head] = true
        }
        head = head.Next
    }

    return false
}
