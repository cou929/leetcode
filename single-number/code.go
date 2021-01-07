import "sort"

func singleNumber(nums []int) int {
	return sol1(nums)
}

// using Sort.sort (quicksort or heapsort)
// time: O(n*log(n)), space: O(1)
func sol1(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i = i + 2 {
		if i+1 < len(nums) && nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

// simple memoize
// time: O(n), space: O(n/2)
func sol2(nums []int) int {
	memo := make(map[int]int)
	for _, n := range nums {
		memo[n]++
	}
	for k, v := range memo {
		if v == 1 {
			return k
		}
	}
	return -1
}
