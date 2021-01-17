// time: O(1), space: O(1)
func reverseBits(num uint32) uint32 {
    cur := num
    res := uint32(0)
    for i := 0; i < 32; i++ {
        if cur & 1 == 1 {
            res++
        }
        cur = cur >> 1
        if i < 32 - 1 {
            res = res << 1
        }
    }
    return res
}
