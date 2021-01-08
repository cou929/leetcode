// time: O(n/2), space: O(1)
func reverseString(s []byte)  {
    for i := 0; i < len(s)/2; i++ {
        swap := s[i]
        s[i] = s[len(s)-i-1]
        s[len(s)-i-1] = swap
    }
}
