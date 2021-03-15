type pos struct {
    x, y int
}

var dirs = []pos{pos{1, 0}, pos{-1, 0}, pos{0, 1}, pos{0, -1}}
var c = 0

func solve(board [][]byte)  {
    vi := make(map[pos]bool)
    for i := 1; i < len(board)-1; i++ {
        for j := 1; j < len(board[i])-1; j++ {
            if vi[pos{i, j}] || board[i][j] == 'X' {
                continue
            }
            m := make([]pos, 0)
            ok := s(i, j, board, vi, &m)
            if !ok {
                continue
            }
            for _, n := range m {
                board[n.x][n.y] = 'X'
            }
        }
    }
    return
}

func s(i, j int, b [][]byte, v map[pos]bool, m *[]pos) bool {
    v[pos{i, j}] = true
    ok := true
    *m = append(*m, pos{i, j})
    for _, d := range dirs {
        ni, nj := i + d.x, j + d.y
        if ni < 0 || nj < 0 || ni >= len(b) || nj >= len(b[0]) {
            continue
        }
        if v[pos{ni, nj}] || b[ni][nj] == 'X' {
            continue
        }
        if ni == 0 || nj == 0 || ni == len(b)-1 || nj == len(b[0])-1 {
            ok = false
            continue
        }
        ok2 := s(ni, nj, b, v, m)
        ok = ok && ok2
    }
    return ok
}
