// time: O(n), space: O(1)
func maxArea(height []int) int {
    return sol2(height)
}

func sol1(height []int) int {
    left := 0
    right := len(height) - 1
    i := 0
    j := len(height) - 1
    res := area(left, right, height)
    for i < j {
        if height[left] < height[right] {
            if height[left] < height[i] {
                left = i
            } else {
                i++
            }
        } else {
            if height[right] < height[j] {
                right = j
            } else {
                j--
            }
        }
        a := area(left, right, height)
        if a > res {
            res = a
        }
    }
    return res
}

func sol2(height []int) int {
    left := 0
    right := len(height) - 1
    res := area(left, right, height)
    for left < right {
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
        a := area(left, right, height)
        if a > res {
            res = a
        }
    }
    return res
}

func area(left, right int, height []int) int {
    h := height[left]
    if h > height[right] {
        h = height[right]
    }
    w := right - left
    return h * w
}
