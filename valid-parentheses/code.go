func isValid(s string) bool {
    return sol1(s)
}

// time: O(n), space: O(n)
func sol1(s string) bool {
    var stack []rune
    for _, r := range s {
        switch r {
            case '{', '(', '[':
            var toStack rune
            switch r {
                case '{':
                toStack = '}'
                case '(':
                toStack = ')'
                case '[':
                toStack = ']'
            }
            stack = append(stack, toStack)
            case '}', ')', ']':
            if len(stack) == 0 {
                return false
            }
            if stack[len(stack)-1] != r {
                return false
            }
            stack = stack[0:len(stack)-1]
        }
    }
    return len(stack) == 0
}
