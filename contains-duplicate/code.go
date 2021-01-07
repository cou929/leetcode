func containsDuplicate(nums []int) bool {
	memo := make(map[int]bool)
	for _, n := range nums {
		if _, ok := memo[n]; ok {
			return true
		}
		memo[n] = true
	}
	return false
}
