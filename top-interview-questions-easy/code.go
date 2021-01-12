/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// time: O(2n), space: O(1)
// search center position of linked-list and reverse first half.
// then compare each element from center iterately.
// argument linked-list is broken by this func. is this correct?
func isPalindrome(head *ListNode) bool {
    count := 0
    counterCursor := head
    revHead := head
    revPrev := (*ListNode)(nil)
    revHeadPos := 0
    
    for counterCursor != nil {
        count++
        counterCursor = counterCursor.Next
        if count/2 > revHeadPos {
            revHeadPos++
            next := revHead.Next
            revHead.Next = revPrev
            revPrev = revHead
            revHead = next
        }
    }

    left := revPrev
    right := revHead
    if count % 2 == 1 {
        right = right.Next
    }
    
    for left != nil && right != nil {
        if left.Val != right.Val {
            return false
        }
        left = left.Next
        right = right.Next
    }

    return true
}
