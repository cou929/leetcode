func lengthOfLongestSubstring(s string) int {
    return sol2(s)
}

// time: O(n), space: O(n)
func sol2(s string) int {
    lastIndex := make(map[rune]int)
    longest := 0
    left := 0
    for right, r := range s {
        if idx, ok := lastIndex[r]; ok && left < idx + 1 {
            left = idx + 1
        }
        if right - left + 1 > longest {
            longest = right - left + 1
        }
        lastIndex[r] = right
    }
    return longest
}

// time: O(n^2), space: O(n)
func sol1(s string) int {
    longest := 0
    l := 0
    memo := make(map[rune]int)
    for i, r := range s {
        if pos, ok := memo[r]; ok {
            if longest < l {
                longest = l
            }
            l = i - pos
            memo = make(map[rune]int)
            for j := pos; j <= i; j++ {
                memo[rune(s[j])] = j
            }
            continue
        }
        l++
        memo[r] = i
    }
    if longest < l {
        longest = l
    }
    return longest
}
