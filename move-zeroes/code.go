// time: O(n), space: O(1)
func moveZeroes(nums []int) {
	cur := 0
	for _, n := range nums {
		if n != 0 {
			nums[cur] = n
			cur++
		}
	}
	for ; cur < len(nums); cur++ {
		nums[cur] = 0
	}
}
