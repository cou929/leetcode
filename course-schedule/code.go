// time: O(n), space:O(n) (n = length of prerequisites)
func canFinish(numCourses int, prerequisites [][]int) bool {
    g := make(map[int][]int)
    ok := make(map[int]bool)
    for _, n :=  range prerequisites {
        g[n[0]] = append(g[n[0]], n[1])
        ok[n[0]] = false
    }
    for k, _ := range g {
        if ok[k] {
            continue
        }
        vi := make(map[int]bool)
        if searchCyclic(k, g, vi, ok) {
            return false
        }
    }
    return true
}

func searchCyclic(cur int, g map[int][]int, vi map[int]bool, ok map[int]bool) bool {
    if vi[cur] {
        return true
    }
    if ok[cur] {
        return false
    }
    ok[cur] = true
    vi[cur] = true
    for _, n := range g[cur] {
        if searchCyclic(n, g, vi, ok) {
            return true
        }
        vi[n] = false
    }
    return false
}
