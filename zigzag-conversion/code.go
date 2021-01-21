// time: O(n); space: O(n)
func convert(s string, numRows int) string {
    slopeLen := numRows - 2
    if slopeLen < 0 {
        slopeLen = 0
    }
    res := []byte{}
    for i := 0; i < numRows; i++ {
        if i == 0 || i == numRows-1 {
            for j := 0; i + (numRows + slopeLen) * j < len(s); j++ {
                res = append(res, s[i + (numRows + slopeLen) * j])
            }
            continue
        }
        
        for j := 0; ; j++ {
            idx := i + (numRows + slopeLen) * j
            if idx >= len(s) {
                break
            }
            res = append(res, s[idx])
            
            idx = numRows * (j + 1) + slopeLen * j + (numRows-1-i) - 1
            if idx >= len(s) {
                break
            }
            res = append(res, s[idx])
        }
    }
    return string(res)
}
