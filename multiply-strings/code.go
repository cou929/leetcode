// example:
//   123
//   456
//   ---
//   738
//  615  =  6888
// 492   = 56088
// -----
// 
func multiply(num1 string, num2 string) string {
    if num1[0] == '0' || num2[0] == '0' {
        return "0"
    }
    res := []int{0}
    for i := 0; i < len(num2); i++ {
        k := int(num2[len(num2)-i-1] - 48)
        num3 := make([]int, i, len(num1)+i)
        carry := 0
        for j := 0; j < len(num1); j++ {
            l := int(num1[len(num1)-j-1] - 48)
            m := k * l + carry
            n := m % 10
            carry = m / 10
            num3 = append([]int{n}, num3...)
        }
        if carry > 0 {
            num3 = append([]int{carry}, num3...)
        }
        carry = 0
        for i := 0; i < len(num3); i++ {
            n3 := num3[len(num3)-i-1]
            n1 := 0
            if len(res)-i-1 >= 0 {
                n1 = res[len(res)-i-1]
            }
            m := n1 + n3 + carry
            n := m % 10
            carry = m / 10
            if len(res)-i-1 >= 0 {
                res[len(res)-i-1] = n
            } else {
                res = append([]int{n}, res...)
            }
        }
        if carry > 0 {
            res = append([]int{carry}, res...)
        }
    }
    res1 := ""
    for i := len(res)-1; i >= 0; i-- {
        res1 = fmt.Sprint(res[i]) + res1
    }
    return res1
}
