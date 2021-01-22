/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    return sol2(root)
}

type N2 struct {
    Node *TreeNode
    Visited bool
}

func sol2(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var q []*N2
    var res []int
    q = append(q, &N2{root, false})
    for len(q) > 0 {
        n := q[0]
        q = q[1:len(q)]
        
        if n.Visited {
            res = append(res, n.Node.Val)
            continue
        }

        // prepend
        if n.Node.Right != nil {
            q = append([]*N2{&N2{n.Node.Right, false}}, q...)
        }
        q = append([]*N2{&N2{n.Node, true}}, q...)
        if n.Node.Left != nil {
            q = append([]*N2{&N2{n.Node.Left, false}}, q...)
        }
    }
    return res
}

func sol1(node *TreeNode) []int {
    if node == nil {
        return nil
    }
    res := []int{}
    if left := sol1(node.Left); left != nil {
        res = append(res, left...)
    }
    res = append(res, node.Val)
    if right := sol1(node.Right); right != nil {
        res = append(res, right...)
    }
    return res
}
