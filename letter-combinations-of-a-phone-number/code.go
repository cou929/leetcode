var charByD map[rune][]string = map[rune][]string{
    '2': []string{"a", "b", "c"},
    '3': []string{"d", "e", "f"},
    '4': []string{"g", "h", "i"},
    '5': []string{"j", "k", "l"},
    '6': []string{"m", "n", "o"},
    '7': []string{"p", "q", "r", "s"},
    '8': []string{"t", "u", "v"},
    '9': []string{"w", "x", "y", "z"},
}

// time: O(4^n), space: O(4^n*2)
func letterCombinations(digits string) []string {
    res := make([]string, 0)
    for i, d := range digits {
        if i == 0 {
            res = append(res, charByD[d]...)
            continue
        }
        l := len(res)
        for _, r := range res {
            for _, s := range charByD[d] {
                res = append(res, r + s)
            }
        }
        res = res[l:len(res)]
    }
    return res
}
