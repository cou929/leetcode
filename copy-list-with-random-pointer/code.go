// time: o(2n), space: o(2n)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    copied := make(map[*Node]*Node)
    res :=  &Node{}
    cur := head
    curRes := res
    for cur != nil {
        c := &Node{cur.Val, nil, nil}
        curRes.Next = c
        copied[cur] = c
        cur = cur.Next
        curRes = curRes.Next
    }
    cur = head
    for cur != nil {
        copied[cur].Random = copied[cur.Random]
        cur = cur.Next
    }
    return res.Next
}
