/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time: o(n), space: o(n) and breaking input
func convertBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    cal(root, 0)
    return root
}

func cal(n *TreeNode, sum int) int {
    s := sum
    if n.Right != nil {
        s = cal(n.Right, s)
    }
    s += n.Val
    n.Val = s
    if n.Left != nil {
        s = cal(n.Left, s)
    }
    return s
}
