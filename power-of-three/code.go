func isPowerOfThree(n int) bool {
    return sol1(n)
}

func sol1(n int) bool {
    if n < 1 {
        return false
    }
    cur := n
    for cur > 0 {
        if cur % 3 != 0 && cur != 1 {
            return false
        }
        cur = cur / 3
    }
    return true
}
