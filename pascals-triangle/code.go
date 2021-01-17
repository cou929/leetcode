// time: O(n*n/2), space: o(n*n/2)
func generate(numRows int) [][]int {
    var res [][]int
    if numRows == 0 {
        return res
    }
    res = append(res, []int{1})
    for i := 1; i < numRows; i++ {
        row := make([]int, len(res[i-1])+1)
        for j := 0; j < len(res[i-1])+1; j++ {
            col := 0
            if j - 1 >= 0 {
                col = col + res[i-1][j-1]
            }
            if j < len(res[i-1]) {
                col = col + res[i-1][j]
            }
            row[j] = col
        }
        res = append(res, row)
    }
    return res
}
