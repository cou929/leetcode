// time: O(n), space: O(1)
func canJump(nums []int) bool {
    need := 0
    for i := len(nums)-2; i >= 0; i-- {
        need++
        if nums[i] >= need {
            need = 0
            continue
        }
    }
    if need > 0 {
        return false
    }
    return true
}
