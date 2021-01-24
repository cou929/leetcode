// time: o(n), space: o(1)
func canCompleteCircuit(gas []int, cost []int) int {
    cand := -1
    sum := 0
    rem := 0
    for i := 0; i < len(gas); i++ {
        sum = sum + gas[i] - cost[i]
        if cand != -1 {
            rem = rem + gas[i] - cost[i]
            if rem < 0 {
                cand = -1
                rem = 0
            }
        } else {
            if gas[i] - cost[i] >= 0 {
                cand = i
                rem = gas[i] - cost[i]
            }
        }
    }
    if sum < 0 {
        return -1
    }
    return cand
}
