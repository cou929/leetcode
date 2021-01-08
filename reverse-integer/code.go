// time: log10(n), space: log(1)
func reverse(xx int) int {
    x := int32(xx)
    res := int32(0)
    for x != 0 {
        if res > 2147483647/10 || (res == 2147483647/10 && x % 10 > 7) {
            return 0
        }
        if res < -2147483648/10 || (res == -2147483648/10 && x % 10 > 8) {
            return 0
        }
        res = res * 10 + x % 10
        x = x / 10
    }
    return int(res)
}
