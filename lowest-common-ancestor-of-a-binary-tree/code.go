// time: o(n), space: o(n)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var pa, qa []*TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    pa, qa = nil, nil
    return s(root, p.Val, q.Val, make([]*TreeNode, 0))
}

func s(n *TreeNode, p, q int, anc []*TreeNode) *TreeNode {
    anc = append(anc, n)

    if n.Val == p {
        pa = make([]*TreeNode, len(anc))
        copy(pa, anc)
    }
    if n.Val == q {
        qa = make([]*TreeNode, len(anc))
        copy(qa, anc)
    }
    if pa != nil && qa != nil {
        return ans()
    }

    if n.Left != nil {
        res := s(n.Left, p, q, anc)
        if res != nil {
            return res
        }
    }
    if n.Right != nil {
        res := s(n.Right, p, q, anc)
        if res != nil {
            return res
        }
    }

    return nil
}

func ans() *TreeNode {
    for i := len(pa)-1; i >= 0; i-- {
        for j := len(qa)-1; j >= 0; j-- {
            if pa[i].Val == qa[j].Val {
                return pa[i]
            }
        }
    }
    return &TreeNode{}
}
