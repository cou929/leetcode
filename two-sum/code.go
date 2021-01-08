func twoSum(nums []int, target int) []int {
	return sol2(nums, target)
}

// brute-force
// time: O(n^2), space: O(1)
func sol1(nums []int, target int) []int {
	for i, n := range nums {
		for j, m := range nums {
			if i == j {
				continue
			}
			if n+m == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// hash map
// time: O(2n), space: O(n)
func sol2(nums []int, target int) []int {
	memo := make(map[int]int)
	for i, n := range nums {
		memo[n] = i
	}
	for i, n := range nums {
		if _, ok := memo[target-n]; ok && i != memo[target-n] {
			return []int{i, memo[target-n]}
		}
	}
	return nil
}
