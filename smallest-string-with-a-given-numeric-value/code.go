// time: O(n), space: O(1)
func getSmallestString(n int, k int) string き{
    res := make([]byte, n)
    rest := k
    for i:= 1; i <= n; i++ {
        t := rest - (n - i)
        if t >= 26 {
            res[n-i] = 'z'
            rest -= 26
            continue
        }
        res[n-i] = byte('a' + t - 1)
        rest -= t
    }
    return string(res)
}
