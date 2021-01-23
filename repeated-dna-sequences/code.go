// time: O(n), space: O(4^10)
func findRepeatedDnaSequences(s string) []string {
    if len(s) < 10 {
        return nil
    }
    memo := make(map[string]int)
    for i := 0; i + 10 <= len(s); i++ {
        cur := s[i:i+10]
        memo[cur]++
    }
    var res []string
    for k, v := range memo {
        if v > 1 {
            res = append(res, k)
        }
    }
    return res
}
