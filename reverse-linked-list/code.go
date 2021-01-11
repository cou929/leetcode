/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    return sol1(head)
}

// iterative
// time: O(n), space: O(1)
// n is a length of the LinkedList
func sol1(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    revNext := (*ListNode)(nil)
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = revNext
        revNext = cur
        cur = next
    }
    return revNext
}

// recursive
// time: O(n), space: O(n) (call stack = n)
func sol2(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    return rev(nil, head)
}

func rev(cur *ListNode, next *ListNode) *ListNode {
    if next == nil {
        return cur
    }
    nextNext := next.Next
    next.Next = cur
    return rev(next, nextNext)
}
