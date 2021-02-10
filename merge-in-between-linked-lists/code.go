/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// time: o(n+m), space: o(1)
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    var left1, right1, head2, tail2 *ListNode
    i := 0
    cur := list1
    for cur != nil {
        if i == a - 1 {
            left1 = cur
        }
        if i == b + 1 {
            right1 = cur
        }
        i++
        cur = cur.Next
    }
    head2 = list2
    cur = list2
    for cur.Next != nil {
        cur = cur.Next
    }
    tail2 = cur
    left1.Next = head2
    tail2.Next = right1
    return list1
}
