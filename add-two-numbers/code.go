/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// time: o(n), space: o(1)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    a := l1
    b := l2
    head := &ListNode{}
    cur := head
    for true {
        v := carry
        if a != nil {
            v = v + a.Val
            a = a.Next
        }
        if b != nil {
            v = v + b.Val
            b = b.Next
        }
        carry = v / 10
        v = v % 10
        cur.Val = v
        if a == nil && b == nil {
            break
        }
        cur.Next = &ListNode{}
        cur = cur.Next
    }
    if carry > 0 {
        cur.Next = &ListNode{}
        cur.Next.Val = carry
    }
    return head
}
