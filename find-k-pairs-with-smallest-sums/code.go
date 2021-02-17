func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    return sol2(nums1, nums2, k)
}

// greedy 2 (like heap solution but not use heap)
// time: o(n1 * n2 * n1), space: o(n1)
func sol2(nums1 []int, nums2 []int, k int) [][]int {
    if len(nums1) == 0 || len(nums2) == 0 {
        return nil
    }

    cursors := make([]int, len(nums1))
    var res [][]int
    for len(res) < k && len(res) < len(nums1) * len(nums2) {
        min := nums1[len(nums1)-1] + nums2[len(nums2)-1] + 1
        next := -1
        for i, j := range cursors {
            if j >= len(nums2) {
                continue
            }
            t := nums1[i] + nums2[j]
            if min > t {
                min = t
                next = i
            }
            // todo: skip condition
        }
        res = append(res, []int{nums1[next], nums2[cursors[next]]})
        cursors[next]++
    }

    return res
}

// greedy (mistake)
func sol1(nums1 []int, nums2 []int, k int) [][]int {
    if len(nums1) == 0 || len(nums2) == 0 {
        return nil
    }
    var res, memo [][]int
    for _, n := range nums1 {
        memo = append(memo, []int{n, nums2[0]})
    }
    res = append(res, memo[0])
    if len(res) == k {
        return res
    }
    midx := 1
    i := 0
    mmax := 0
    for ; i < len(nums1); i++ {
            j := 1
        fmt.Println("s", i, j, res)
        if midx < len(memo) {
            mmax = memo[midx][0] + memo[midx][1]
        }
        for ; j < len(nums2) && nums1[i] + nums2[j] < mmax; j++ {
            res = append(res, []int{nums1[i], nums2[j]})
        }
        if midx < len(memo) {
            res = append(res, []int{memo[midx][0], memo[midx][1]})
        }
        midx++
        if midx >= len(memo) {
            for ; j < len(nums2); j++ {
                res = append(res, []int{nums1[i], nums2[j]})
            }
        }
        fmt.Println("e", i, j, res)
        if len(res) >= k {
            break
        }
    }
    if len(res) > k {
        res = res[0:k]
    }
    return res
}
