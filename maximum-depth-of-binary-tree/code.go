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
    Depth int
}

// time: O(n), space: O(n)
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    res := 0
    lifo := []*Elm{}
    lifo = append([]*Elm{&Elm{root, 1}}, lifo...)

    for len(lifo) > 0 {
        n := lifo[0]
        lifo = lifo[1:len(lifo)]

        if n.Depth > res {
            res = n.Depth
        }
        
        if n.Node.Right != nil {
            lifo = append([]*Elm{&Elm{n.Node.Right, n.Depth+1}}, lifo...)
        }
        if n.Node.Left != nil {
            lifo = append([]*Elm{&Elm{n.Node.Left, n.Depth+1}}, lifo...)
        }
    }

    return res
}
