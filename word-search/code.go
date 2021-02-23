func exist(board [][]byte, word string) bool {
    return sol2(board, word)
}

// complexity is same to sol1
// but does not copy map every loop
func sol2(board [][]byte, word string) bool {
    visited := make(map[pos]bool)
    for x := 0; x < len(board); x++ {
        for y := 0; y < len(board[0]); y++ {
            if board[x][y] == word[0] {
                if ok := check(x, y, board, word, visited); ok {
                    return true
                }
            }
        }
    }
    return false
}

func check(x, y int, board [][]byte, rest string, visited map[pos]bool) bool {
    if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
        return false
    }
    if visited[pos{x, y}] {
        return false
    }
    if board[x][y] != rest[0] {
        return false
    }
    if len(rest) == 1 {
        return true
    }
    shifts := []pos{pos{1, 0}, pos{0, 1}, pos{-1, 0}, pos{0, -1}}
    visited[pos{x, y}] = true
    for _, s := range shifts {
        if ok := check(x + s.x, y + s.y, board, rest[1:len(rest)], visited); ok {
            return true
        }
    }
    visited[pos{x, y}] = false
    return false
}

// time: O(n * m * word.length), space: O((n * m)^word.length) ?
func sol1(board [][]byte, word string) bool {
    q := make([]que, 0, 100)
    for x := 0; x < len(board); x++ {
        for y := 0; y < len(board[0]); y++ {
            if board[x][y] == word[0] {
                q = append(q, que{1, pos{x, y}, map[pos]bool{pos{x, y}: true}})
            }
        }
    }
    shifts := []pos{pos{1, 0}, pos{0, 1}, pos{-1, 0}, pos{0, -1}}
    for len(q) > 0 {
        cur := q[len(q)-1]
        q = q[0:len(q)-1]
        if cur.validLen == len(word) {
            return true
        }
        for _, s := range shifts {
            x := cur.pos.x + s.x
            y := cur.pos.y + s.y
            if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
                continue
            }
            if board[x][y] != word[cur.validLen] || cur.visited[pos{x, y}] {
                continue
            }
            tmpVisited := copyMap(cur.visited)
            tmpVisited[pos{x, y}] = true
            q = append(q, que{cur.validLen + 1, pos{x, y}, tmpVisited})
        }
    }
    return false
}

type que struct {
    validLen int
    pos pos
    visited map[pos]bool
}

type pos struct {
    x, y int
}

func copyMap(in map[pos]bool) map[pos]bool {
    res := make(map[pos]bool)
    for k, v := range in {
        res[k] = v
    }
    return res
}
