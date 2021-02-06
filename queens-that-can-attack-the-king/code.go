type pos struct {
    x int
    y int
}

// time: o(n), space: o(n)
func queensAttacktheKing(queens [][]int, king []int) [][]int {
    var res [][]int
    q := make(map[pos]bool)
    for _, v := range queens {
        q[pos{v[0], v[1]}] = true
    }
    // up
    for i := 1; ; i++ {
        if king[0] - i < 0 {
            break
        }
        x := king[0] - i
        y := king[1]
        if _, ok := q[pos{x, y}]; ok {
            res = append(res, []int{x, y})
            break
        }        
    }
    // up-right
    for i := 1; ; i++ {
        if king[0] - i < 0  || king[1] + i >= 8 {
            break
        }
        x := king[0] - i
        y := king[1] + i
        if _, ok := q[pos{x, y}]; ok {
            res = append(res, []int{x, y})
            break
        }        
    }
    // right
    for i := 1; ; i++ {
        if king[1] + i >= 8 {
            break
        }
        x := king[0]
        y := king[1] + i
        if _, ok := q[pos{x, y}]; ok {
            res = append(res, []int{x, y})
            break
        }        
    }
    // down-right
    for i := 1; ; i++ {
        if king[0] + i >= 8 || king[1] + i >= 8 {
            break
        }
        x := king[0] + i
        y := king[1] + i
        if _, ok := q[pos{x, y}]; ok {
            res = append(res, []int{x, y})
            break
        }        
    }
    // down
    for i := 1; ; i++ {
        if king[0] + i >= 8 {
            break
        }
        x := king[0] + i
        y := king[1]
        if _, ok := q[pos{x, y}]; ok {
            res = append(res, []int{x, y})
            break
        }        
    }
    // down-left
    for i := 1; ; i++ {
        if king[0] + i >= 8 || king[1] - i < 0 {
            break
        }
        x := king[0] + i
        y := king[1] - i
        if _, ok := q[pos{x, y}]; ok {
            res = append(res, []int{x, y})
            break
        }        
    }
    // left
    for i := 1; ; i++ {
        if king[1] - i < 0 {
            break
        }
        x := king[0]
        y := king[1] - i
        if _, ok := q[pos{x, y}]; ok {
            res = append(res, []int{x, y})
            break
        }        
    }
    // up-left
    for i := 1; ; i++ {
        if king[0] - i < 0 || king[1] - i < 0 {
            break
        }
        x := king[0] - i
        y := king[1] - i
        if _, ok := q[pos{x, y}]; ok {
            res = append(res, []int{x, y})
            break
        }        
    }
    return res
}
