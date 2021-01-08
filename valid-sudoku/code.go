import "strconv"

// time: O(1), space: O(1)
func isValidSudoku(board [][]byte) bool {
	memo1 := make(map[int]bool)
	memo2 := make(map[int]bool)

	// row and col
	for i := 0; i < len(board); i++ {
		memo1 = make(map[int]bool)
		memo2 = make(map[int]bool)
		for j := 0; j < len(board); j++ {
			// row
			v, _ := strconv.Atoi(string(board[i][j]))
			if _, ok := memo1[v]; ok && v != 0 {
				return false
			}
			memo1[v] = true

			// col
			v, _ = strconv.Atoi(string(board[j][i]))
			if _, ok := memo2[v]; ok && v != 0 {
				return false
			}
			memo2[v] = true
		}
	}

	// sub-box
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			baseX := i * 3
			baseY := j * 3
			memo1 = make(map[int]bool)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					v, _ := strconv.Atoi(string(board[baseX+k][baseY+l]))
					if _, ok := memo1[v]; ok && v != 0 {
						return false
					}
					memo1[v] = true
				}
			}
		}
	}

	return true
}
