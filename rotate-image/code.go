// time: O(n^2), space: O(1)
func rotate(matrix [][]int)  {
    for phase := 0; phase < len(matrix)/2; phase++ {
        l := len(matrix) - phase * 2
        bases := [][]int{
            []int{0 + phase, 0 + phase},
            []int{0 + phase, len(matrix) - phase - 1},
            []int{len(matrix) - phase - 1, len(matrix) - phase - 1},
            []int{len(matrix) - phase - 1, 0 + phase},
        }
        next := get(matrix, bases[0][0], bases[0][1], 0, l-1)
        var swap []int
        for i := 0; i < len(bases); i++ {
            destIdx := (i + 1) % len(bases)
            swap = get(matrix, bases[destIdx][0], bases[destIdx][1], destIdx, l-1)
            put(matrix, bases[destIdx][0], bases[destIdx][1], destIdx, next)
            next = swap
        }
    }
}

// dir 0:right, 1:down, 2:left, 3:up
func get(matrix [][]int, baseX int, baseY int, dir int, l int) []int {
    res := make([]int, l)
    for i := 0; i < l; i++ {
        var x, y int
        switch dir {
        case 0:
            x = baseX
            y = baseY + i
        case 1:
            x = baseX + i
            y = baseY
        case 2:
            x = baseX
            y = baseY - i
        case 3:
            x = baseX - i
            y = baseY
        }
        res[i] = matrix[x][y]
    }
    return res
}

// dir 0:right, 1:down, 2:left, 3:up
func put(matrix [][]int, baseX int, baseY int, dir int, nums []int) {
    for i := 0; i < len(nums); i++ {
        var x, y int
        switch dir {
        case 0:
            x = baseX
            y = baseY + i
        case 1:
            x = baseX + i
            y = baseY
        case 2:
            x = baseX
            y = baseY - i
        case 3:
            x = baseX - i
            y = baseY
        }
        matrix[x][y] = nums[i]
    }
}
