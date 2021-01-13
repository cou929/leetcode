/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time: O(n), space: O(n)
func levelOrder(root *TreeNode) [][]int {
    var res [][]int
    if root == nil {
        return res
    }
    q := []*TreeNode{root}
    for len(q) > 0 {
        var curLevel []int
        l := len(q)
        for i := 0; i < l; i++ {
            n := q[0]
            q = q[1:len(q)]
            curLevel = append(curLevel, n.Val)
            if n.Left != nil {
                q = append(q, n.Left)
            }
            if n.Right != nil {
                q = append(q, n.Right)
            }
        }
        res = append(res, curLevel)
    }
    return res
}
