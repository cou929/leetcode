/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time: o(n), space: o(n) (call  stack)
func sumNumbers(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return s(root, 0)
}

func s(n *TreeNode, ss int) int {
    res := 0
    sum := ss * 10 + n.Val
    if n.Left == nil && n.Right == nil {
        return sum
    }
    if n.Left != nil {
        res += s(n.Left, sum)
    }
    if n.Right != nil {
        res += s(n.Right, sum)
    }
    return res
}
