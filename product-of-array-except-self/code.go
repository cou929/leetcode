func productExceptSelf(nums []int) []int {
    return sol2(nums)
}

// time: o(2n), space: o(1)
func sol2(nums []int) []int {
    res := make([]int, len(nums))
    // left
    res[0] = 1
    for i, _ := range nums {
        if i == 0 {
            continue
        }
        res[i] = res[i-1] * nums[i-1]
    }
    // right
    r := 1
    for i, _ := range nums {
        if i == 0 {
            continue
        }
        r *= nums[len(nums)-i-1+1]
        res[len(nums)-i-1] *= r
    }
    return res
}

// time: o(2n), space: o(1)
func sol1(nums []int) []int {
    p := 1
    z := 0
    for _, n := range nums {
        if n == 0 {
            z++
            if z >= 2 {
                break
            }
            continue
        }
        p *= n
    }
    if z >= 2 {
        p = 0
    }
    res := make([]int, len(nums))
    for i, n := range nums {
        if z == 1 {
            res[i] = 0
            if n == 0 {
                res[i] = p
            }
            continue
        }
        if n == 0 {
            res[i] = 0
            continue
        }
        res[i] = p / n
    }
    return res
}
