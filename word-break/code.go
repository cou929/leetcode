func wordBreak(s string, wordDict []string) bool {
    d := make(map[string]bool)
    for _, s := range wordDict {
        d[s] = true
    }
    m := make(map[string]bool)
    return b(s, d, m)
}

func b(s string, d map[string]bool, m map[string]bool) bool {
    if res, ok := m[s]; ok {
        return res
    }
    for i := 0; i < len(s); i++ {
        c := s[0:len(s)-i]
        if d[c] {
            rest := s[len(s)-i:len(s)]
            if len(rest) == 0 || b(rest, d, m) {
                m[s] = true
                return true
            }
        }
    }
    m[s] = false
    return false
}
