func removeDuplicates(nums []int) int {
	var cur int
	pos := 0

	for i, n := range nums {
		if i <= 0 {
			cur = n
		}
		if n != cur {
			pos++
			nums[pos] = n
			cur = n
		}
	}

	if pos > len(nums) {
		nums = nums[0 : pos+1]
	}

	return pos + 1
}
