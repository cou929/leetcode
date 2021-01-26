/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// time: o(n+m), space: o(n+m)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    a := []int{}
    n:= l1
    for n != nil {
        a = append(a, n.Val)
        n = n.Next
    }
    b := []int{}
    n = l2
    for n != nil {
        b = append(b, n.Val)
        n = n.Next
    }
    carry := 0
    n = nil
    for i := 0; i < len(a) || i < len(b); i++ {
        ai := len(a) - 1 - i
        bi := len(b) - 1 - i
        t := 0
        if ai >= 0 {
            t = a[ai]
        }
        if bi >= 0 {
            t += b[bi]
        }
        t += carry
        carry = 0
        if t >= 10 {
            carry = 1
        }
        nn := &ListNode{t % 10, n}
        n = nn
    }
    if carry > 0 {
        nn := &ListNode{carry, n}
        n = nn
    }
    return n
}
