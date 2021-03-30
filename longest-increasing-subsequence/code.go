// time: o(n^2), space: o(n)
func lengthOfLIS(nums []int) int {
    dp := make([]int, len(nums))
    res := 1
    for i, n := range nums {
        max := 1
        for j := 0; j < i; j++ {
            if nums[j] < n && max < dp[j] + 1 {
                max = dp[j] + 1
            }
        }
        dp[i] = max
        if res < max {
            res = max
        }
    }
    return res
}
