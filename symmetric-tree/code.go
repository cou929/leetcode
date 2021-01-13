/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time: O(n), space: O(2^log(n)) ?
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        for i := 0; i <= len(queue)/2; i++ {
            l := queue[i].Left
            r := queue[len(queue)-i-1].Right
            if (l == nil) != (r == nil) {
                return false
            }
            if l != nil && l.Val != r.Val {
                return false
            }
            
            l = queue[i].Right
            r = queue[len(queue)-i-1].Left
            if (l == nil) != (r == nil) {
                return false
            }
            if l != nil && l.Val != r.Val {
                return false
            }

        }
        l := len(queue)
        for i := 0; i < l; i++ {
            node := queue[0]
            queue = queue[1:len(queue)]
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    return true
}
