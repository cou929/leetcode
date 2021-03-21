type n struct {
    p []int
    c []int
}

type dep struct {
    v, d int
}

func findOrder(nc int, prs [][]int) []int {
    g := make(map[int]*n)
    for _, p := range prs {
        if _, ok := g[p[0]]; !ok {
            g[p[0]] = &n{make([]int, 0), make([]int, 0)}
        }
        if _, ok := g[p[1]]; !ok {
            g[p[1]] = &n{make([]int, 0), make([]int, 0)}
        }
        g[p[0]].c = append(g[p[0]].c, p[1])
        g[p[1]].p = append(g[p[1]].p, p[0])
    }
    d := make([]dep, nc)
    for i, _ := range d {
        d[i] = dep{i, -1}
    }
    vi := make(map[int]bool)
    for k, v := range g {
        if len(v.p) == 0 {
            if ok := dfs(k, 0, g, vi, d); !ok {
                return nil
            }
        }
    }
    c := 0
    for _, dd := range d {
        if dd.d == -1 {
            c++
        }
    }
    if nc - len(g) < c {
        return nil
    }
    sort.Slice(d, func(i, j int) bool { return d[i].d > d[j].d })
    res := make([]int, nc)
    for i, _ := range d {
        res[i] = d[i].v
    }
    return res
}

func dfs(k, curd int, g map[int]*n, vi map[int]bool, d []dep) bool {
    if vi[k] {
        return false
    }
    if d[k].d < curd {
        d[k].d = curd
    }
    vi[k] = true
    for _, next := range g[k].c {
        if ok := dfs(next, curd+1, g, vi, d); !ok {
            return false
        }
    }
    vi[k] = false
    return true
}
