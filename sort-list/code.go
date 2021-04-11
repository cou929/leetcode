/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    return mergeSort(head)
}

// time: o(nlogn), space: o(1)
func mergeSort(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    return ms(head, nil)
}

func ms(left, right *ListNode) *ListNode {
    if left == right || left.Next == right {
        return &ListNode{left.Val, nil}
    }
    n := 1
    mn := 0
    mid := left
    cur := left
    for cur.Next != right {
        cur = cur.Next
        n++
        if n/2 > mn {
            mid = mid.Next
            mn++
        }
    }
    a := ms(left, mid)
    b := ms(mid, right)
    return merge(a, b)
}

func merge(a, b *ListNode) *ListNode {
    head := &ListNode{0, nil}
    curA, curB, curRes := a, b, head
    for curA != nil || curB != nil {
        if curA != nil && curB != nil {
            if curA.Val < curB.Val {
                curRes.Next = &ListNode{curA.Val, nil}
                curRes = curRes.Next
                curA = curA.Next
            } else {
                curRes.Next = &ListNode{curB.Val, nil}
                curRes = curRes.Next
                curB = curB.Next                
            }
        } else if curA != nil {
            curRes.Next = &ListNode{curA.Val, nil}
            curRes = curRes.Next
            curA = curA.Next
        } else {
            curRes.Next = &ListNode{curB.Val, nil}
            curRes = curRes.Next
            curB = curB.Next
        }
    }
    return head.Next
}

// time: o(n^2), space: o(1)
func liner(head *ListNode) *ListNode {
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
