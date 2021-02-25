func threeSum(nums []int) [][]int {
    return sol2(nums)
}

// time: O(n^2), space: O(n)
// breaks a input
func sol2(nums []int) [][]int {
    sort.Ints(nums)
    memo := make(map[int]int)
    for i, n := range nums {
        memo[n] = i
    }
    res := make([][]int, 0)
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            break
        }
        if i > 0 && nums[i-1] == nums[i] {
            continue
        }
        for j := i+1; j < len(nums); j++ {
            if j > i+1 && nums[j-1] == nums[j] {
                continue
            }
            n, m := nums[i], nums[j]
            if memo[-(n+m)] <= j {
                continue
            }
            res = append(res, []int{n, m, -(n+m)})
        }
    }
    return res
}

// wip
func sol1(nums []int) [][]int {
    memo := make(map[int][]int)
    for i, n := range nums {
        if memo[n] == nil {
            memo[n] = make([]int, 0, 1)
        }
        memo[n] = append(memo[n], i)
    }
    res := make([][]int, 0)
    for i := 0; i < len(nums); i ++ {
        for j := i + 1; j < len(nums); j++ {
            n, m := nums[i], nums[j]
            if len(memo[-(n+m)]) <= 0{
                continue
            }
            cands := memo[-(n+m)]
            for _, k := range cands {
                if k > j {
                    res = append(res, []int{n, m, nums[k]})
                }
            }
        }
    }
    return res
}
