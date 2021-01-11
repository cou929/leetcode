// time: O(n), space: O(1) n is length of linked list
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    node := head
    c := 0
    cand := head
    for true {
        c++
        if c - n - 1 > 0 {
            cand = cand.Next
        }
        if node.Next == nil {
            break
        }
        node = node.Next
    }
    if c - n == 0 {
        head = head.Next
        return head
    }
    cand.Next = cand.Next.Next
    return head
}
