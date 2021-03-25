// time: o(nlogn), space: o(n)

type t struct {
    n, f int
}

func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    for _, n := range nums {
        m[n]++
    }
    s := make([]t, 0, len(m))
    for k, v := range m {
        s = append(s, t{k, v})
    }
    sort.Slice(s, func(i, j int) bool { return s[i].f > s[j].f })
    res := make([]int, k)
    for i := 0; i < k; i++ {
        res[i] = s[i].n
    }
    return res
}
