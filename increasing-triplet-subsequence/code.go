// time: O(n) (3n), space: O(n) (2n)
func increasingTriplet(nums []int) bool {
    if len(nums) < 3 {
        return false
    }
    minLeft := make([]int, len(nums))
    maxRight := make([]int, len(nums))
    minLeft[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        if minLeft[i-1] < nums[i] {
            minLeft[i] = minLeft[i-1]
        } else {
            minLeft[i] = nums[i]
        }
    }
    maxRight[len(nums)-1] = nums[len(nums)-1]
    for i := len(nums)-2; i >= 0; i-- {
        if maxRight[i+1] > nums[i] {
            maxRight[i] = maxRight[i+1]
        } else {
            maxRight[i] = nums[i]
        }
    }
    for i := 1; i < len(nums)-1; i++ {
        if minLeft[i-1] < nums[i] && nums[i] < maxRight[i+1] {
            return true
        }
    }
    return false
}
