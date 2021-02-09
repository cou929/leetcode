// time: o(n), space: o(n)
func partitionDisjoint(A []int) int {
    mins := make([]int, len(A))
    m := 1000001
    for i, _ := range A {
        idx := len(A) - 1 - i
        if m > A[idx] {
            m = A[idx]
        }
        mins[idx] = m
    }
    m = 0
    for i := 0; i < len(A) - 1; i++ {
        if m < A[i] {
            m = A[i]
        }
        if m <= mins[i+1] {
            return i + 1
        }
    }
    return -1
}
