func rotate(nums []int, k int)  {
    sol3(nums, k)
}

// time: O(n), space: O(n)
func sol1(nums []int, k int) {
    pos := len(nums) - (k % len(nums))
    for i, n := range append(nums[pos:len(nums)], nums[0:pos]...) {
        nums[i] = n
    }
}

// time: O(n), space: O(1)
func sol2(nums []int, k int) {
    if len(nums) == 0 || k == 0 {
        return
    }
    
    next := nums[0] // todo
    swap := 0
    cur := 0
    done := 0
    
    for done < len(nums) {
        done++

        dest := (cur + k) % len(nums)
        //fmt.Println(cur, dest, next)

        swap = nums[dest]
        nums[dest] = next
        
        unit := k
        if (len(nums) - k < unit && len(nums) - k > 0) {
            unit = len(nums) - k
        }
        if len(nums) % unit == 0 && done % (len(nums) / unit) == 0 {
            dest = (dest + 1) % len(nums)
            swap = nums[dest]
        }

        cur = dest
        next = swap
    }
}

// brute-force
// time: O(N), space: O(1)
func sol3(nums []int, k int) {
    c := k % len(nums)
    for i := 0; i < c; i++ {
        right(nums)
    }
}

func right(nums []int) {
    swap := nums[len(nums)-1]
    for i := len(nums)-1; i > 0; i-- {
        nums[i] = nums[i-1]
        
    }
    nums[0] = swap
}
