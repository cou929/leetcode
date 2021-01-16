// time: O(n), space: O(n)
func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    dp := make([]int, len(nums))
    for i, n := range nums {
        a := 0
        if i - 2 >= 0 {
            a = dp[i-2]
        }
        b := 0
        if i - 1 >= 0 {
            b = dp[i-1]
        }
        dp[i] = b
        if a + n > b {
            dp[i] = a + n
        }
    }
    return dp[len(dp)-1]
}
