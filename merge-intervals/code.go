// time: O(nlogn + n), space: O(logn)
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
    res := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        cur := intervals[i]
        last := res[len(res)-1]
        if last[1] >= cur[0] {
            if last[1] < cur[1] {
                last[1] = cur[1]
            }
            continue
        }
        res = append(res, cur)
    }
    return res
}
