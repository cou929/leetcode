// time: O(n * 2^n), space: O(n * 2^n)
func subsets(nums []int) [][]int {
    res := make([][]int, 0)
    res = append(res, []int{})
    for _, n := range nums {
        for _, r := range res {
            v := make([]int, len(r), len(r) + 1)
            copy(v, r)
            v = append(v, n)
            res = append(res, v)
        }
    }
    return res
}
