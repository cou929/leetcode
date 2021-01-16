func maxProfit(prices []int) int {
    min := prices[0]
    max := -1
    p := -1
    
    for i := 1; i < len(prices); i++ {
        if min > prices[i] {
            min = prices[i]
            max = -1
        } else if max < prices[i]  {
            max = prices[i]
        }
        if max - min > p {
            p = max - min
        }
    }
    if p == -1 {
        return 0
    }
    return p
}
