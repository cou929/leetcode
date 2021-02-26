// time: O(n*m*(n+m) + n*m), space: O(1)
func setZeroes(matrix [][]int)  {
    rand.Seed(time.Now().UnixNano())
    var toZero int
    for {
        toZero = rand.Int()
        for i, _ := range matrix {
            for _, n := range matrix[i] {
                if n == toZero {
                    goto CONT
                }
            }
        }
        break
        CONT:
    }
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            if matrix[i][j] == 0 {
                for k := 0; k < len(matrix); k++ {
                    if matrix[k][j] != 0 {
                        matrix[k][j] = toZero
                    }
                }
                for k := 0; k < len(matrix[0]); k++ {
                    if matrix[i][k] != 0 {
                        matrix[i][k] = toZero
                    }
                }
            }
        }
    }
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            if matrix[i][j] == toZero {
                matrix[i][j] = 0
            }
        }
    }
    return
}
