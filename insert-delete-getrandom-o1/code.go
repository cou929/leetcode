type RandomizedSet struct {
    h map[int]int
    a []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    rand.Seed(time.Now().UnixNano())
    return RandomizedSet{
        h: make(map[int]int),
        a: make([]int, 0, 100001),
    }
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.h[val]; ok {
        return false
    }
    this.h[val] = len(this.a)
    this.a = append(this.a, val)
    return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.h[val]; !ok {
        return false
    }
    if this.a[len(this.a)-1] != val {
        pos := this.h[val]
        this.a[pos] = this.a[len(this.a)-1]
        this.h[this.a[pos]] = pos
    }
    delete(this.h, val)
    this.a = this.a[0:len(this.a)-1]
    return true
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    i := rand.Intn(len(this.a))
    return this.a[i]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
