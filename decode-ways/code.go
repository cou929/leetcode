// time: o(n), space: o(1)
func numDecodings(s string) int {
    nums := make([]int, 101)
    nums[1], nums[2] = 1, 2
    for i := 3; i <= 100; i++ {
        nums[i] = nums[i-1] + nums[i-2]
    }
    res := 1
    l := 0
    for i, ss := range s {
        cur := ss - '0'
        l++
        if cur == 1 || cur == 2 {
            continue
        }
        if cur == 0 {
            if l == 1 {
                return 0
            }
            if l > 3 {
                res *= nums[l-2]
            } else {
                res *= 1
            }
        } else if cur >= 7 && (l > 1 && s[i-1] - '0' == 2) {
            res *= nums[l-1]
        } else if l > 0 {
            res *= nums[l]
        }
        l = 0
    }
    if l > 0 {
        res *= nums[l]
    }
    return res
}
