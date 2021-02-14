type RandomizedCollection struct {
    h map[int][]int
    a []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
    rand.Seed(time.Now().UnixNano())
    return RandomizedCollection{
        h: make(map[int][]int),
        a: make([]int, 0, 100000),
    }
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
    res := true
    if _, ok := this.h[val]; ok {
        res = false
    }
    this.h[val] = append(this.h[val], len(this.a))
    this.a = append(this.a, val)
    return res
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
    //fmt.Println(this.h, this.a, val)
    if _, ok := this.h[val]; !ok {
        return false
    }
    // delete from location map
    pos := this.h[val][len(this.h[val])-1]
    this.h[val] = this.h[val][0:len(this.h[val])-1]
    if len(this.h[val]) == 0 {
        delete(this.h, val)
    }
    // swap if val to be removed is not last element
    if pos != len(this.a) - 1 {
        last := this.a[len(this.a)-1]
        this.a[pos] = last
        this.h[last][len(this.h[last])-1] = pos
        sort.Ints(this.h[last])  // todo
    }
    // delete swapped value
    this.a = this.a[0:len(this.a)-1]
    //fmt.Println("removed", this.h, this.a)
    return true
}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
    return this.a[rand.Intn(len(this.a))]
}


/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
