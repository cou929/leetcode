type key struct {
	pos      int
	isBought bool
}

var memo map[key]int

func maxProfit(prices []int) int {
	memo = make(map[key]int)
	return try(prices, 0, -1)
}

func try(prices []int, pos int, bought int) int {
	if m, ok := memo[key{pos, bought != -1}]; ok {
		if bought != -1 {
			return m - bought
		} else {
			return m
		}
	}

	if pos+1 == len(prices) {
		if bought != -1 {
			memo[key{pos, true}] = prices[pos]
			return prices[pos] - bought
		}
		memo[key{pos, false}] = 0
		return 0
	}

	if bought == -1 {
		a := try(prices, pos+1, prices[pos])
		b := try(prices, pos+1, -1)
		if a > b {
			memo[key{pos, false}] = a
			return a
		}
		memo[key{pos, false}] = b
		return b
	} else {
		a := (prices[pos] - bought) + try(prices, pos+1, -1)
		b := try(prices, pos+1, bought)
		if a > b {
			memo[key{pos, true}] = a + bought
			return a
		}
		memo[key{pos, true}] = b + bought
		return b
	}
}
