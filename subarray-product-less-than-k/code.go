// time: o(2n), space: o(1)
func numSubarrayProductLessThanK(nums []int, k int) int {
    // [10, 5, 2, 6, 1000, 2] k == 100
    // 10 => 10
    // 50 => 10-5, 5
    // 100 -> 10 => 5-2, 2
    // 60 => 5-2-6, 2-6, 6
    // 60000 -> 12000 -> 6000 -> 1000 => n/a
    // 2000 -> 2 => 2
    if k <= 1 {
        return 0
    }
    res := 0
    left := 0
    right := 0
    m := 1
    for right < len(nums) {
        m *= nums[right]
        for m >= k && left <= right {
            m /= nums[left]
            left++
        }
        if m < k {
            res += right - left + 1
        }
        right++
    }
    return res
}
