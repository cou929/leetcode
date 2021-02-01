func checkSubarraySum(nums []int, k int) bool {
    return sol2(nums, k)
}

// dp using prefix-sum (total sum mod k)
// time: o(n), space: o(n)
func sol2(nums []int, k int) bool {
    psmod := make(map[int]int)
    psmod[0] = -1 // index 0 (psmod では -1) からの累計が mod k == 0 になるケースのための番兵
    s := 0
    for i, n := range nums {
        s += n
        if k != 0 {  // k == 0 の場合 ps[j] == s (j < i-1)が成り立つもの (つまり連続した 0) を探せば良い。psmod には単に累計を入れれば良い
            s = s % k
        }
        if j, ok := psmod[s]; ok {  // 累計に対する mod k が同じということは、その間に k の倍数だけ増加しているということ。これが探したい答え
            if i - j > 1 {  // psmod[i] には nums[i] を組み入れた値が入っているので、i - j が 2 以上離れていないといけない
                return true
            }
        } else {  // すでに同値の psmod があり、それと 2 以上離れていない場合は更新しない。遠い方 (より左の方) で解消できればそれでよいし、更新するとその可能性をカバーできない
            psmod[s] = i
        }
    }
    return false
}

// prefix-sum & blute-force
func sol1(nums []int, k int) bool {
    ps := make(map[int]int)
    ps[-1] = 0
    for i, n := range nums {
        ps[i] = ps[i-1] + n
    }
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            s := ps[j] - ps[i-1]
            if k == 0 {
                if s == 0 {
                    return true
                }
                continue
            }
            if s % k == 0 {
                return true
            }
        }
    }
    return false
}
