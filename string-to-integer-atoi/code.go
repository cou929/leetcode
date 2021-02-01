// time: O(n), space: O(1)
func myAtoi(s string) int {
    isNeg := false
    firstChar := false
    res := 0
    for _, r := range s {
        if !firstChar && unicode.IsSpace(r) {
            continue
        }
        if !firstChar {
            if !unicode.IsDigit(r) && r != '-' && r != '+' {
                return 0
            }
            firstChar = true
            if r == '-' {
                isNeg = true
                continue
            }
            if r == '+' {
                continue
            }
        }
        if !unicode.IsDigit(r) {
            break
        }
        digit := int(r) - 48
        if isNeg && (res > 2147483648/10 || (res == 2147483648/10 && digit > 8)) {
            return -2147483648
        }
        if !isNeg && (res > 2147483647/10 || (res == 2147483647/10 && digit > 7)) {
            return 2147483647
        }
        res = (res * 10) + digit
    }
    if isNeg {
        return -res
    }
    return res
}
