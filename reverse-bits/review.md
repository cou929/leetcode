# https://leetcode.com/problems/reverse-bits/

- 単純な 1 bit ごとの走査で AC
- solution より
    - 二分探索的に左右をスワップしていく
    - 対象範囲をマスクでとりだして、シフトさせ、or をとるとループなしでいける
    - bit 演算のコツとしての側面（マスク活用）とreverse 実装の側面（半々に分割統治）が勉強になった

```go
func reverseBits(num uint32) uint32 {
    num = (num >> 16) | (num << 16)
    num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
    num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
    num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
    num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
    return num
}
```
