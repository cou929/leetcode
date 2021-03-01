func search(nums []int, target int) int {
    for i, n := range nums {
        if n == target {
            return i
        }
    }
    return -1
}
