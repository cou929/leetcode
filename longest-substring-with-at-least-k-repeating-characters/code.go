func longestSubstring(s string, k int) int {
    return sol2(s, k)
}

// recursive
// time: o(2n * (num of recursion)), space: o(num of recursion)
// num of recursion = approx n in worst case
// so time: o(n^2), space: o(n)  ???
func sol2(s string, k int) int {
    base := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        base[s[i]]++
    }
    notOk := make(map[byte]bool)
    for b, v := range base {
        if v < k {
            notOk[b] = true
        }
    }
    if len(notOk) == 0 {
        return len(s)
    }
    res := 0
    left := 0
    right := 0
    for ; right < len(s); right++ {
        if notOk[s[right]] {
            if right - left >= k {
                l := sol2(s[left:right], k)
                if l > res {
                    res = l
                }
            }
            left = right + 1
        }
    }
    if right - left >= k {
        l := sol2(s[left:right], k)
        if l > res {
            res = l
        }
    }
    return res
}

// brute force (tle)
// time: o(26 * n^2), space: o(1) ???
func sol1(s string, k int) int {
    //fmt.Println(s)
    for i:= 0; i <= len(s) - k; i++ {
        l := len(s) - i
        for j := 0; j + l <= len(s); j++ {
            m := make(map[byte]int)
            for n := 0; n < l; n++ {
                m[s[j + n]]++
            }
            if ok(m, k) {
                // fmt.Println(m)
                return l
            }
        }
    }
    return 0
}

func ok(in map[byte]int, k int) bool {
    if len(in) == 0 {
        return false
    }
    for _, v := range in {
        if v < k {
            return false
        }
    }
    return true
}
