func permute(nums []int) [][]int {
    return sol1(nums)
}

type x struct {
    res []int
    rest []int
}

// time: O(n!), space: O(n!)
func sol1(nums []int) [][]int {
    if len(nums) == 1 {
        return [][]int{nums}
    }
    var result [][]int
    q := make([]x, 0)
    for i, n := range nums {
        rest := make([]int, len(nums))
        copy(rest, nums)
        rest = append(rest[0:i], rest[i+1:len(nums)]...)
        q = append(q, x{[]int{n}, rest})
    }
    for len(q) > 0 {
        cur := q[0]
        q = q[1:len(q)]
        for i, n := range cur.rest {
            t := cur

            res := make([]int, len(t.res))
            copy(res, t.res)
            res = append(res, n)
            t.res = res

            rest := make([]int, len(t.rest))
            copy(rest, t.rest)
            rest = append(rest[0:i], rest[i+1:len(rest)]...)
            t.rest = rest

            if len(t.res) == len(nums) {
                result = append(result, t.res)
                continue
            }
            q = append(q, t)
        }
    }
    return result
}
