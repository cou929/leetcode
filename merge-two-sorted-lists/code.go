/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// time: O(n+m), space: O(1)
// n and m are length of l1 and l2
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var head, node, cur1, cur2 *ListNode
    cur1 = l1
    cur2 = l2

    if cur1 == nil && cur2 == nil {
        return cur1
    }
    if cur1 == nil {
        return cur2
    }
    if cur2 == nil {
        return cur1
    }
    if cur1.Val <= cur2.Val {
        head = cur1
        cur1 = cur1.Next
    } else {
        head = cur2
        cur2 = cur2.Next
    }
    node = head

    for cur1 != nil || cur2 != nil {
        if cur1 == nil || (cur2 != nil && cur2.Val < cur1.Val) {
            node.Next = cur2
            node = node.Next
            cur2 = cur2.Next
            continue
        }
        if cur2 == nil || (cur1 != nil && cur1.Val <= cur2.Val) {
            node.Next = cur1
            node = node.Next
            cur1 = cur1.Next
            continue
        }
    }
    
    return head
}
