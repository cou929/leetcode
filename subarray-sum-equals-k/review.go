// time: o(n), space: o(n)
func subarraySum(nums []int, k int) int {
    m := make(map[int]int)
    m[0] = 1
    res := 0
    sum := 0
    for _, n := range nums {
        sum += n
        if i, ok := m[sum-k]; ok {
            res += i
        }
        m[sum] += 1
    }
    return res
}
