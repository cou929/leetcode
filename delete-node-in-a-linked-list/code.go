/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// time: O(1), space: O(1)
func deleteNode(node *ListNode) {
    nextnext := node.Next.Next
    node.Val = node.Next.Val
    node.Next = nextnext
}
