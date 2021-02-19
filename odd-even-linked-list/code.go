/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// time: o(n), space: o(1)
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    c := 0
    oddLast := head
    evenHead := head.Next
    cur := head
    for cur != nil {
        c++
        if c % 2 == 1 {
            oddLast = cur
        }
        next := cur.Next
        if next != nil {
            cur.Next = next.Next
        }
        cur = next
    }
    oddLast.Next = evenHead
    return head
}
