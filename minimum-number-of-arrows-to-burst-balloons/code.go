// time: o(n*log(n)+n), space: o(1)
// n*log(n) is average case of quicksort (Go's sort package)
func findMinArrowShots(points [][]int) int {
    if len(points) == 0 {
        return 0
    }
    sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })
    res := 1
    r := points[0][1]
    for _, n := range points {
        if r >= n[0] {
            if n[1] < r {
                r = n[1]
            }
            continue
        }
        res++
        r = n[1]
    } 
    return res
}
