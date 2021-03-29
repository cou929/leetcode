// time: o(n), space: o(n)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    stack := []*Node{root}
    next := (*Node)(nil)
    for len(stack) > 0 {
        nStack := make([]*Node, 0)
        for _, n := range stack {
            n.Next = next
            next = n
            if n.Right != nil {
                nStack = append(nStack, n.Right)
            }
            if n.Left != nil {
                nStack = append(nStack, n.Left)
            }
        }
        next = nil
        stack = nStack
    }
    return root
}
