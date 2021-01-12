/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    return sol2(head)
}

// time: O(n), space: O(1)
// it's easy to solve with breaking input list, but is this correct?
func sol1(head *ListNode) bool {
    cur := head
    mark := 100001
    for cur != nil {
        if cur.Val == mark {
            return true
        }
        cur.Val = mark
        cur = cur.Next
    }
    return false
}

// Floyd's Cycle Finding Algorithm from solution
// time: O(n), space: O(1)
func sol2(head *ListNode) bool {
    if head == nil {
        return false
    }
    slow := head
    fast := head.Next
    for fast != nil && fast.Next != nil {
        if slow == fast {
            return true
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    return false
}
