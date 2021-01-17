// time: O(n) (n = num of digits of x, y in binary), space: O(1)
func hammingDistance(x int, y int) int {
    diff := x ^ y
    res := 0
    for diff > 0 {
        if diff & 0b1 == 0b1 {
            res++
        }
        diff = diff >> 1
    }
    return res
}
