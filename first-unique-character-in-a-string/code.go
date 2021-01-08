// time: O(2n), space: O(n)
func firstUniqChar(s string) int {
    memo := make(map[rune]int)
    for _, c := range s {
        memo[c]++
    }
    for i, c := range s {
        if memo[c] == 1 {
            return i
        }
    }
    return -1
}
