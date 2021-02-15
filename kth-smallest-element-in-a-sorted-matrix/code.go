func kthSmallest(matrix [][]int, k int) int {
    return sol1(matrix, k)
}

func sol1(matrix [][]int, k int) int {
    cursors := make([]int, len(matrix))
    res := -1
    for i := 0; i < k; i++ {
        next := 0
        min := matrix[len(matrix)-1][len(matrix)-1]
        for j, n := range cursors {
            if n >= len(matrix) {
                continue
            }
            if min > matrix[j][n] {
                min = matrix[j][n]
                next = j
            }
        }
        res = min
        cursors[next]++
    }
    return res
}
