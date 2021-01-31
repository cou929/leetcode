// time: O(2n + ?), space: O(n)
type c struct {
    Left int
    Right int
}

type memo struct {
    cs []c
    best int
}

const big = 100000 + 1 // 10^5 + 1

func minSumOfLengths(arr []int, target int) int {
    m := &memo{
        best: big,
    }
    left := 0
    right := 0
    sum := arr[right]
    for left < len(arr) && right < len(arr) {
        if sum == target {
            update(left, right, m)
        }
        if sum < target {
            right++
            if right < len(arr) {
                sum += arr[right]
            }
        } else {
            sum -= arr[left]
            left++
            if left > right {
                right++
                if right < len(arr) {
                    sum += arr[right]
                }
            }
        }
    }
    if m.best == big {
        return -1
    }
    return m.best
}

// - 見つけた候補について、過去と照らし、最適を超えなかったら push しない
// - かぶっている候補について、直前の候補と同じ長さなら push しない (長い場合はそっちがベストの可能性があるので push)
func update(left, right int, m *memo) {
    if len(m.cs) == 0 {
        m.cs = append(m.cs, c{left, right})
        return
    }
    last := m.cs[len(m.cs)-1]
    if last.Right >= left && (last.Right - last.Left) == (right - left) {
        return
    }

    m.cs = append(m.cs, c{left, right})
    best := findBest(m.cs)
    switch {
    case best == -1:
        return
    case best < m.best:
        m.best = best
        return
    default:
        m.cs = m.cs[0:len(m.cs)-1]
    }
    return
}

// brute-force
func findBest(m []c) int {
    res := big
    for i := 0; i < len(m); i++ {
        for j := i + 1; j < len(m); j++ {
            if m[i].Right >= m[j].Left {
                continue
            }
            l := (m[i].Right - m[i].Left + 1) + (m[j].Right - m[j].Left + 1)
            if res > l {
                res = l
            }
        }
    }
    if res == big {
        res = -1
    }
    return res
}
