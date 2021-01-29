// time: o(1), space: o(1)
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
    area := (C - A) * (D - B) + (G - E) * (H - F)
    switch {
         case C < E, G < A, D < F, H < B:
        return area
    }
    left := A
    if A < E {
        left = E
    }
    right := C
    if G < C {
        right = G
    }
    top := D
    if H < D {
        top = H
    }
    bottom := B
    if B < F {
        bottom = F
    }
    return area - (right - left) * (top - bottom)
}
