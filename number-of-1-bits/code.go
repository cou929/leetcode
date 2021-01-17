func hammingWeight(num uint32) int {
    return sol2(num)
}

func sol1(num uint32) int {
    res := 0
    cur := num
    for cur > 0 {
        if cur % 2 == 1 {
            res++
        }
        cur = cur / 2
    }
    return res
}

func sol2(num uint32) int {
    res := 0
    cur := num
    for cur > 0 {
        if cur & 0b1 == 0b1 {
            res++
        }
        cur = cur >> 1
    }
    return res
}
