// time: O(n) (O(100n)), space: O(n)
func numSquares(n int) int {
    psq := make([]int, 0, 100)
    for i := 1; i * i <= 10000; i++ {
        psq = append(psq, i*i)
    }
    m := make([]int, n+1)
    for i := 1; i <= n; i++ {
        step := m[i-psq[0]] + 1
        for j := 1; j < len(psq); j++ {
            if psq[j] > i {
                break
            }
            cand := m[i-psq[j]] + 1
            if step > cand {
                step = cand
            }
        }
        m[i] = step
    }
    return m[n]
}
