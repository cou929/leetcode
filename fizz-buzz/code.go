func fizzBuzz(n int) []string {
    var res []string
    for i := 1; i <= n; i++ {
        s := strconv.Itoa(i)
        if i % 3 == 0 {
            s = "Fizz"
        }
        if i % 5 == 0 {
            s = "Buzz"
        }
        if i % 3 == 0 && i % 5 == 0 {
            s = "FizzBuzz"
        }
        res = append(res, s)
    }
    return res
}
