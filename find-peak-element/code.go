func findPeakElement(nums []int) int {
    return sol2(nums)
}

// bisect
// time: o(logn), space: o(1)
func sol2(nums []int) int {
    left, right, mid := 0, len(nums)-1, (len(nums)-1)/2
    for left < right {
        if isPeak(mid, nums) {
            break
        }
        l, r := nums[mid], nums[mid]
        if mid > 0 {
            l = nums[mid-1]
        }
        if mid < len(nums)-1 {
            r = nums[mid+1]
        }
        if l < r {
            left = mid
        } else {
            right = mid
        }
        if left == mid && left + 1 == right {
            mid = right
            continue
        }
        mid = (right - left) / 2 + left
        // 2mid = right - left + 2left
        // 2mid = right + left
        // mid = (right + left) / 2
    }
    return mid
}

func isPeak(i int, nums []int) bool {
    if i > 0 && nums[i-1] >= nums[i] {
        return false
    }
    if i < len(nums) - 1 && nums[i+1] >= nums[i] {
        return false
    }
    return true
}

// time: O(n), space: O(1)
func sol1(nums []int) int {
    for i, n := range nums {
        if i > 0 && nums[i-1] >= n {
            continue
        }
        if i < len(nums) - 1 && nums[i+1] >= n {
            continue
        }
        return i
    }
    return -1
}
