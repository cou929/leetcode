import "math/rand"

type Solution struct {
    orig []int
}


func Constructor(nums []int) Solution {
    return Solution{
        orig: nums,
    }    
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.orig    
}


/** Returns a random shuffling of the array. */
// time: O(n), space: O(n) (for orig)
func (this *Solution) Shuffle() []int {
    n := make([]int, len(this.orig))
    copy(n, this.orig)
    for i, _ := range n {
        j := rand.Int() % len(n)
        n[i], n[j] = n[j], n[i]
    }
    return n
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
