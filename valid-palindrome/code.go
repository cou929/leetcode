func isPalindrome(s string) bool {
    cleaned := []rune{}
    for _, r := range s {
        if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
            continue
        }
        cleaned = append(cleaned, unicode.ToLower(r))
    }
    for i := 0; i < len(cleaned)/2; i++ {
        if cleaned[i] != cleaned[len(cleaned)-1-i] {
            return false
        }
    }
    return true
}
