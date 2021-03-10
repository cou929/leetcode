// time:o(nm), space:o(nm)
func numIslands(grid [][]byte) int {
    is := make([][]int, len(grid))
    for i := 0; i < len(grid); i++ {
        is[i] = make([]int, len(grid[0]))
    }
    c := 1
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if is[i][j] != 0 {
                continue
            }
            search(i, j, grid, is)
            if grid[i][j] == '1' {
                c++
            }
        }
    }
    return c - 1
}

func search(i, j int, grid [][]byte, is [][]int) {
    if grid[i][j] == '0' {
        is[i][j] = -1
        return
    }
    ds := [][]int{
        []int{1, 0}, []int{-1, 0},
        []int{0, 1}, []int{0, -1},
    }
    is[i][j] = 1
    for _, d := range ds {
        ni := i + d[0]
        nj := j + d[1]
        if ni < 0 || nj < 0 || ni >= len(grid) || nj >= len(grid[0]) {
            continue
        }
        if is[ni][nj] != 0 {
            continue
        }
        search(ni, nj, grid, is)
    }
}
