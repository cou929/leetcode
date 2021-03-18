type dll struct {
    Key int
    Val int
    Next *dll
    Prev *dll
}

func (d *dll) Disconnect() {
    d.Prev.Next = d.Next
    d.Next.Prev = d.Prev
}

func (d *dll) AddNext(in *dll) {
    n := d.Next
    d.Next = in
    in.Next = n
    in.Prev = d
    n.Prev = in
}

type LRUCache struct {
    capa int
    key2dll map[int]*dll
    head, tail *dll
}


func Constructor(capacity int) LRUCache {
    head, tail := &dll{}, &dll{}
    head.Next = tail
    tail.Prev = head
    return LRUCache{
        capa: capacity,
        key2dll: make(map[int]*dll),
        head: head,
        tail: tail,
    }
}


func (this *LRUCache) Get(key int) int {
    d, ok := this.key2dll[key]
    if !ok {
        return -1
    }
    d.Disconnect()
    this.head.AddNext(d)
    return d.Val
}


func (this *LRUCache) Put(key int, value int)  {
    d, ok := this.key2dll[key]
    if ok {
        d.Val = value
        d.Disconnect()
        this.head.AddNext(d)
        return
    }
    d = &dll{key, value, nil, nil}
    if len(this.key2dll) >= this.capa {
        t := this.tail.Prev
        t.Disconnect()
        delete(this.key2dll, t.Key)
        t = nil
    }
    this.head.AddNext(d)
    this.key2dll[key] = d
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
