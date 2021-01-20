// time: O(n*n/2), space: O(1)
func longestPalindrome(s string) string {
    mr := 0
    ml := 0
    for i, _ := range s {
        l := i
        r := i
        for r < len(s) && s[l] == s[r] {
            if (mr - ml) < (r - l) {
                mr = r
                ml = l
            }
            r++
        }
        r--
        for l >= 0 && r < len(s) {
            if s[l] != s[r] {
                break
            }
            if (mr - ml) < (r - l) {
                mr = r
                ml = l
            }
            l--
            r++
        }
    }
    return s[ml:mr+1]
}
