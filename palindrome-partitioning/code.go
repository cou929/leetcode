// time: o(n^2), space: o(?)
func partition(s string) [][]string {
    m := make([][][]string, len(s))
    for i, _ := range s {
        m[i] = make([][]string, 0)
        for j := i; j >= 0; j-- {
            sub := s[j:i+1]
            if isPal(sub) {
                if j-1 >= 0 {
                    for _, ss := range m[j-1] {
                        t := make([]string, len(ss))
                        copy(t, ss)
                        m[i] = append(m[i], append(t, sub))
                    }
                } else {
                    m[i] = append(m[i], []string{sub})
                }
            }
        }
    }
    return m[len(m)-1]
}

func isPal(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-i-1] {
            return false
        }
    }
    return true
}
