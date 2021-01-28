// time: o(n), space: o(n)
func longestValidParentheses(s string) int {
    open := 0
    dp := make([]int, len(s))
    max := 0
    for i, r := range s {
        t := 0
        switch r {
            case '(':
                open++
            case ')':
            switch open {
                case 0:
                    // nothing
                default:
                    t = dp[i-1] + 2
                    if i - t >= 0 {
                        t += dp[i-t]
                    }
                    open--
            }
        }
        dp[i] = t
        if max < t {
            max = t
        }
    }
    return max
}
