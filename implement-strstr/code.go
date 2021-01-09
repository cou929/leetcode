// time: O(nm), space: O(1)
func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }
    if len(needle) > len(haystack) {
        return -1
    }
    for i := 0; i < len(haystack); i++ {
        if haystack[i] == needle[0] {
            j := 0
            for ; j < len(needle) && i + j < len(haystack); j++ {
                if haystack[i+j] != needle[j] {
                    break
                }
            }
            if j == len(needle) {
                return i
            }
        }
    }
    return -1
}
