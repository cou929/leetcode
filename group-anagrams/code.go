// time: O(n * klog(k)), space: O(nk), supporce k = length of each string
func groupAnagrams(strs []string) [][]string {
    var res [][]string
    index := make(map[string]int)
    for _, str := range strs {
        r := []rune(str)
        sort.Sort(sortRunes(r))
        i, ok := index[string(r)]
        if !ok {
            index[string(r)] = len(res)
            res = append(res, []string{str})
            continue
        }
        res[i] = append(res[i], str)
    }
    return res
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}
