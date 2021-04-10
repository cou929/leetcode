// time: o(n^2), space: o(1)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    left := &ListNode{-100001, nil}
    cur := head
    for cur != nil {
        t := left
        for true {
            if t.Next == nil {
                break
            }
            if t.Next.Val > cur.Val {
                break
            }
            t = t.Next
        }
        n := t.Next
        t.Next = &ListNode{cur.Val, n}
        cur = cur.Next
    }
    return left.Next
}
