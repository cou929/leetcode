// brute-force
// time: O(nm) m?, space: O(2m)
func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    
    nums := []int{1}
    for i := 2; i <= n; i++ {
        cur := -1
        count := 0
        swap := []int{}
        for _, m := range nums {
            if cur != m {
                if cur != -1 {
                    swap = append(swap, count, cur)
                }
                cur = m
                count = 0
            }
            count++
        }
        swap = append(swap, count, cur)
        nums = swap
    }
    res := ""
    for _, v := range nums {
        res = res + strconv.Itoa(v)
    }
    return res
}
