type MyCircularDeque struct {
    Head *Node
    Size int
    Len int
}

type Node struct {
    Val int
    Next *Node
    Prev *Node
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        Head: nil,
        Size: k,
        Len: 0,
    }
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }    
    n := &Node{
        Val: value,
    }
    if this.IsEmpty() {
        this.Head = n
        this.Head.Next = this.Head
        this.Head.Prev = this.Head
    } else {
        n.Next = this.Head
        n.Prev = this.Head.Prev
        this.Head.Prev = n
        this.Head = n   
    }
    this.Len++
    return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }    
    n := &Node{
        Val: value,
    }
    if this.IsEmpty() {
        this.Head = n
        this.Head.Next = this.Head
        this.Head.Prev = this.Head
    } else {
        n.Next = this.Head
        n.Prev = this.Head.Prev
        this.Head.Prev.Next = n
        this.Head.Prev = n   
    }
    this.Len++
    return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    p := this.Head.Prev
    this.Head = this.Head.Next
    this.Head.Prev = p
    this.Len--
    return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    newLast := this.Head.Prev.Prev
    newLast.Next = this.Head
    this.Head.Prev = newLast
    this.Len--
    return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.Head.Val
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.Head.Prev.Val
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    return this.Len == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    return this.Len == this.Size
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
