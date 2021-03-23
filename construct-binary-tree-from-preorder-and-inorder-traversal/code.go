// time:o(n), space:o(1)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(pre []int, in []int) *TreeNode {
    root := &TreeNode{pre[0], nil, nil}
    div := -1
    for i, n := range in {
        if n == pre[0] {
            div = i
            break
        }
    }
    if len(in[0:div]) > 0 {
        in2 := in[0:div]
        pre2 := pre[1:len(in2)+1]
        left := buildTree(pre2, in2)
        root.Left = left
    }
    if len(in[div+1:len(in)]) > 0 {
        in2 := in[div+1:len(in)]
        pre2 := pre[len(pre)-len(in2):len(pre)]
        right := buildTree(pre2, in2)
        root.Right = right
    }
    return root
}
