import "strings"

// time: O(n*log(n)), space: O(1)
func isAnagram(s string, t string) bool {
    ss := strings.Split(s, "")
    sort.Strings(ss)
    tt := strings.Split(t, "")
    sort.Strings(tt)
    return strings.Join(ss, "") == strings.Join(tt, "")
}
