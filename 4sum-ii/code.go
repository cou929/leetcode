// time: o(n * n + n), space: o(n*n * 2)
func fourSumCount(A []int, B []int, C []int, D []int) int {
    return sol2(A, B, C, D)
}

// time: o(n*n * 2), space: o(n*n)
func sol2(A []int, B []int, C []int, D []int) int {
    m := make(map[int]int)
    for i := 0; i < len(A); i++ {
        for j := 0; j < len(A); j++ {
            m[A[i] + B[j]]++
        }
    }
    res := 0
    for i := 0; i < len(A); i++ {
        for j := 0; j < len(A); j++ {
            res += m[-(C[i] + D[j])]
        }
    }
    return res
}

func sol1(A []int, B []int, C []int, D []int) int {
    first := make(map[int]int)
    second := make(map[int]int)
    for i := 0; i < len(A); i++ {
        for j := 0; j < len(A); j++ {
            first[A[i] + B[j]]++
            second[C[i] + D[j]]++
        }
    }
    res := 0
    for n, freq1 := range first {
        if freq2, ok := second[-n]; ok {
            res += freq1 * freq2
        }
    }
    return res
}
