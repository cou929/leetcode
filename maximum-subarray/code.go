// time: O(n), space: O(1)
func maxSubArray(nums []int) int {
    p := -100001
    cur := 0
    for _, n := range nums {
        if cur + n < n {
            cur = n
        } else {
            cur = cur + n
        }
        if cur > p {
            p = cur
        }
    }
    return p
}
