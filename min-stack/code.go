type MinStack struct {
    s []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{[]int{}}    
}


func (this *MinStack) Push(x int)  {
    this.s = append(this.s, x)    
}


func (this *MinStack) Pop()  {
    this.s = this.s[0:len(this.s)-1]
}


func (this *MinStack) Top() int {
    return this.s[len(this.s)-1]
}


func (this *MinStack) GetMin() int {
    res := 2147483648
    for _, n := range this.s {
        if res > n {
            res = n
        }
    }
    return res
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
