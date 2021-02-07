// time: o(nlog(n)) (quicksort), space: o(n)
func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }
    w1 := make(map[byte]int)
    w2 := make(map[byte]int)
    for i, _ := range word1 {
        w1[word1[i]]++
        w2[word2[i]]++
    }
    var s1, s2 []string
    var i1, i2 []int
    for k, v := range w1 {
        s1 = append(s1, string(k))
        i1 = append(i1, v)
    }
    for k, v := range w2 {
        s2 = append(s2, string(k))
        i2 = append(i2, v)
    }
    sort.Strings(s1)
    sort.Strings(s2)
    sort.Ints(i1)
    sort.Ints(i2)
    if len(s1) != len(s2) {
        return false
    }
    for i, _ := range s1 {
        if s1[i] != s2[i] {
            return false
        }
    }
    if len(i1) != len(i2) {
        return false
    }
    for i, _ := range i1 {
        if i1[i] != i2[i] {
            return false
        }
    }
    return true
}
