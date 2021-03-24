// time: o(nlogn), space:o(n)
func largestNumber(nums []int) string {
    s := make([]string, len(nums))
    for i, n := range nums {
        s[i] = strconv.Itoa(n)
    }
    sort.Slice(s, func(i, j int) bool {
        return s[i] + s[j] > s[j] + s[i]
    })
    if s[0] == "0" {
        return "0"
    }
    return strings.Join(s, "")
}
