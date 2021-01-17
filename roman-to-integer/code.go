var toInt map[string]int = map[string]int{
    "I":  1,
    "IV": 4,
    "V":  5,
    "IX": 9,
    "X":  10,
    "XL": 40,
    "L":  50,
    "XC": 90,
    "C":  100,
    "CD": 400,
    "D":  500,
    "CM": 900,
    "M":  1000,
}

func romanToInt(s string) int {
    res := 0
    for i := 0; i < len(s); i++ {
        if i + 1 < len(s) {
            if x, ok := toInt[string(s[i:i+2])]; ok {
                res = res + x
                i++
                continue
            }
        }
        x := toInt[string(s[i])]
        res = res + x
    }
    return res
}
