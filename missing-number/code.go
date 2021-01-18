// time: O(n), space: O(1)
func missingNumber(nums []int) int {
    want := 0
    sum := 0
    for i, n := range nums {
        want = want + i + 1
        sum = sum + n
    }
    return want - sum
}
