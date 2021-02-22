func generateParenthesis(n int) []string {
    res := []string{"()"}
    for i := 1; i < n; i++ {
        memo := make(map[string]bool)
        l := len(res)
        for _, str := range res {
            for i, _ := range str {
                cand := str[0:i] + "()" + str[i:len(str)]
                if !memo[cand] {
                    memo[cand] = true
                    res = append(res, cand)
                }
            }
        }
        res = res[l:len(res)]
    }
    return res
}
