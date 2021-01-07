// time: O(n), space: O(n)
func plusOne(digits []int) []int {
    var res []int
    carry := 1
    for i := len(digits)-1; i >= 0; i-- {
        n := digits[i]
        d := n + carry
        if d >= 10 {
            carry = 1
            res = append([]int{0}, res...)
            continue
        }
        res = append([]int{d}, res...)
        carry = 0
    }
    if carry == 1 {
        res = append([]int{1}, res...)
    }
    return res
}
