/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Elm struct {
    Node *TreeNode
    Max int
    Min int
}

// time: O(n), space: O(n) (log(n)*2 ?)
func isValidBST(root *TreeNode) bool {
    fifo := []*Elm{}
    fifo = append(fifo, &Elm{root, 2147483647 + 1, -2147483648 - 1})
    for len(fifo) > 0 {
        cur := fifo[0]
        fifo = fifo[1:len(fifo)]
        if cur.Node.Left != nil {
            if cur.Min >= cur.Node.Left.Val || cur.Node.Val <= cur.Node.Left.Val {
                return false
            }
            fifo = append(fifo, &Elm{cur.Node.Left, cur.Node.Val, cur.Min})
        }
        if cur.Node.Right != nil {
            if cur.Node.Val >= cur.Node.Right.Val || cur.Max <= cur.Node.Right.Val {
                return false
            }
            fifo = append(fifo, &Elm{cur.Node.Right, cur.Max, cur.Node.Val})
        }
    }
    return true
}
