// time: O(n), space: O(1)
func myPow(x float64, n int) float64 {
    return sol2(x, n)
}

// time: O(log(n)), space: O(1)
func sol2(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    nn := n
    if n < 0 {
        nn = -n
    }
    res := p(x, nn)
    if n < 0 {
        return 1/res
    }
    return res
}

func p(x float64, n int) float64 {
    if n == 1 {
        return x
    }
    res := p(x, n/2)
    if n % 2 == 0 {
        return res * res
    }
    return res * res * x
}

// time: O(n), space: O(1)
func sol1(x float64, n int) float64 {
    if x == 1 || x == 0 || (x == -1 && n % 2 == 1) || (x == -1 && n % 2 == -1) {
        return x
    }
    if x == -1 {
        return -x
    }
    res := 1.000
    if n < 0 {
        x = 1/x
        n = -n
    }
    for i := 0; i < n; i++ {
        res *= x
    }
    return res
}
