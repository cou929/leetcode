func countPrimes(n int) int {
    return sol1(n)
}

// time: O(?), space: O(n)
func sol1(n int) int {
    if n < 2 {
        return 0
    }
    nums := make([]bool, n)
    nums[0] = true
    nums[1] = true
    for i := 2; i*i < n; i++ {
        if nums[i] == true {
            continue
        }
        for j := i*i; j < n; j += i {
            nums[j] = true
        }
    }
    res := 0
    for _, x := range nums {
        if x == false {
            res++
        }
    }
    return res
}
