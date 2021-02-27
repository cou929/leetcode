func sortColors(nums []int) {
    sol2(nums)
}

// time: O(n), space: O(1)
func sol1(nums []int) {
    left, right := 0, len(nums)-1
    for i := 0; i < len(nums); {
        switch nums[i] {
        case 0:
            for left < len(nums) {
                if nums[left] != 0 {
                    break
                }
                left++
            }
            if left < i {
                swap(nums, left, i)
            }
            i++
        case 1:
            i++
            continue
        case 2:
            for right >= 0 {
                if nums[right] != 2 {
                    break
                }
                right--
            }
            if right > i {
                swap(nums, i, right)
            }
        }
        if i > right {
            break
        }
    }
    return
}

// not sure this algorithm is "in-place"
// time: O(n), space: O(1)
func sol2(nums []int) {
    c := make([]int, 3)
    for _, n := range nums {
        c[n]++
    }
    for i, _ := range nums {
        switch {
        case c[0] > 0:
            nums[i] = 0
            c[0]--
            continue
        case c[1] > 0:
            nums[i] = 1
            c[1]--
            continue
        case c[2] > 0:
            nums[i] = 2
            c[2]--
            continue
        }
    }
    return
}

func swap(nums []int, i, j int) {
    swap := nums[i]
    nums[i] = nums[j]
    nums[j] = swap
    return
}
