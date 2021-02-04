func maxProduct(nums []int) int {
    return sol2(nums)
}

// time: O(n), space: O(1)
func sol2(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    res := nums[0]
    pmax := nums[0]
    pmin := nums[0]
    for i := 1; i < len(nums); i++ {
        n := nums[i]
        c1 := pmax * n
        c2 := pmin * n
        pmax = max(max(n, c1), c2)
        pmin = min(min(n, c1), c2)
        if res < pmax {
            res = pmax
        }
    }
    return res
}

func max(m, n int) int {
    if m > n {
        return m
    } else {
        return n
    }
}

func min(m, n int) int {
    if m > n {
        return n
    } else {
        return m
    }
}

// time: O(n^2), space: O(n)
func sol1(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    c := []int{nums[0]}
    res := nums[0]
    for i := 1; i < len(nums); i++ {
        n := nums[i]
        for j, m := range c {
            c[j] = n * m
            if c[j] > res {
                res = c[j]
            }
        }
        if nums[i-1] <= 0 {
            c = append(c, n)
        }
        if n > res {
            res = n
        }
    }
    return res
}
