/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    _, res := r(root, 0, k)
    return res
}

func r(node *TreeNode, c, k int) (int, int) {
    cc := c
    res := 0
    if node.Left != nil {
        cc, res = r(node.Left, cc, k)
        if cc == k {
            return cc, res
        }
    }
    cc++
    if cc == k {
        return cc, node.Val
    }
    if node.Right != nil {
        cc, res = r(node.Right, cc, k)
        if cc == k {
            return cc, res
        }
    }
    return cc, res
}
