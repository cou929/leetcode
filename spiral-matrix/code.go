// time: O(n), space: O(n)
func spiralOrder(matrix [][]int) []int {
    visited := make([][]int, len(matrix))
    for i := 0; i < len(matrix); i++ {
        visited[i] = make([]int, len(matrix[0]))
    }
    dirs := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
    curDir := 0
    curX, curY := 0, 0
    res := make([]int, 0, len(matrix) + len(matrix[0]))
    for {
        res = append(res, matrix[curX][curY])
        visited[curX][curY] = 1
        nextX, nextY := curX + dirs[curDir][0], curY + dirs[curDir][1]
        if nextX >= len(matrix) || nextX < 0 || nextY >= len(matrix[0]) || nextY < 0 || visited[nextX][nextY] == 1 {
            curDir = (curDir + 1) % 4
            nextX, nextY = curX + dirs[curDir][0], curY + dirs[curDir][1]
        }
        if nextX >= len(matrix) || nextX < 0 || nextY >= len(matrix[0]) || nextY < 0 || visited[nextX][nextY] == 1 {
            break
        }
        curX, curY = nextX, nextY
    }
    return res
}
