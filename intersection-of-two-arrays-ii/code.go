// time: O(n*log(n), space: O(1)
func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    nums := nums1
    check := nums2
    if len(nums1) < len(nums2) {
        nums = nums2
        check = nums1
    }
    
    var res []int
    var j = 0
    for i := 0; i < len(nums); {
        n := nums[i]
        if j >= len(check) {
            break
        }
        if n < check[j] {
            i++
            continue
        }
        if n > check[j] {
            j++
            continue
        }
        if n == check[j] {
            res = append(res, n)
            i++
            j++
        }
    }
    
    return res
}
