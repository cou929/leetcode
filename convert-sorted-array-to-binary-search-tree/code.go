/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time: O(n), space: O(1)
type Arg struct {
    Head *TreeNode
    Nums []int
}

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    head := &TreeNode{-1, nil, nil}
    q := []*Arg{&Arg{head, nums}}
    for len(q) > 0 {
        cur := q[0]
        q = q[1:len(q)]
        center := len(cur.Nums) / 2
        cur.Head.Val = cur.Nums[center]
        left := cur.Nums[0:center]
        right := cur.Nums[center+1:len(cur.Nums)]
        if len(left) > 0 {
            cur.Head.Left = &TreeNode{-1, nil, nil}
            q = append(q, &Arg{cur.Head.Left, left})
        }
        if len(right) > 0 {
            cur.Head.Right = &TreeNode{-1, nil, nil}
            q = append(q, &Arg{cur.Head.Right, right})
        }
    }
    return head
}
