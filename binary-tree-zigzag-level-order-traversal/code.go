/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type E struct {
    n *TreeNode
    level int
    right bool
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    q := make([]*E, 0)
    res := make([][]int, 0)
    q = append(q, &E{root, 0, true})
    for len(q) > 0 {
        cur := q[0]
        q = q[1:len(q)]
        if cur.level >= len(res) {
            res = append(res, make([]int, 0, 1))
        }
        res[cur.level] = append(res[cur.level], cur.n.Val)
        i := 0
        for ; i < len(q); i++ {
            if q[i].level > cur.level {
                break
            }
        }
        head := q[0:i]
        rest := q[i:len(q)]
        if cur.right {
            if cur.n.Left != nil {
                rest = append([]*E{&E{cur.n.Left, cur.level+1, !cur.right}}, rest...)
            }
            if cur.n.Right != nil {
                rest = append([]*E{&E{cur.n.Right, cur.level+1, !cur.right}}, rest...)
            }
        } else {
            if cur.n.Right != nil {
                rest = append([]*E{&E{cur.n.Right, cur.level+1, !cur.right}}, rest...)
            }
            if cur.n.Left != nil {
                rest = append([]*E{&E{cur.n.Left, cur.level+1, !cur.right}}, rest...)
            }
        }
        q = append(head, rest...)
    }
    return res
}
