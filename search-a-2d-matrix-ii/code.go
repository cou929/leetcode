// time: o(nlogn), space: o(1)
type win struct {
    sx, sy, ex, ey int
}

func searchMatrix(matrix [][]int, target int) bool {
    return s(matrix, target, win{0, 0, len(matrix)-1, len(matrix[0])-1})
}

func s(mx [][]int, targ int, w win) bool {
    if mx[w.sx][w.sy] > targ {
        return false
    }
    if mx[w.ex][w.ey] < targ {
        return false
    }
    anchorx, anchory := w.sx, w.sy
    l := w.ex - w.sx
    if w.ey - w.sy < l {
        l = w.ey - w.sy
    }
    for i := 0; i <= l; i++ {
        x, y := w.sx + i, w.sy + i
        if mx[x][y] == targ {
            return true
        }
        if mx[x][y] > targ {
            break
        }
        anchorx, anchory = x, y
    }
    if anchory+1 <= w.ey {
        if ok := s(mx, targ, win{w.sx, anchory+1, anchorx, w.ey}); ok {
            return true
        }
    }
    if anchorx + 1 <= w.ex {
        if ok := s(mx, targ, win{anchorx+1, w.sy, w.ex, anchory}); ok {
            return true
        }
    }
    return false
