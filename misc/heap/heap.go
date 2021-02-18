package heap

import "fmt"

// MinHeap is MinHeap data structure.
// Supports int only
type MinHeap struct {
	Tree []int
}

// NewMinHeap is constructor of MinHeap
func NewMinHeap(size int) *MinHeap {
	return &MinHeap{
		Tree: make([]int, 0, size),
	}
}

func (h *MinHeap) len() int {
	return len(h.Tree)
}

func (h *MinHeap) downLeft(cur int) int {
	next := 2*cur + 1
	if next >= h.len() {
		return -1
	}
	return next
}

func (h *MinHeap) downRight(cur int) int {
	next := 2*cur + 2
	if next >= h.len() {
		return -1
	}
	return next
}

func (h *MinHeap) up(cur int) int {
	if cur == 0 {
		return -1
	}
	return (cur - 1) / 2
}

func (h *MinHeap) heapifyUp(pos int) {
	fmt.Println(h.Tree)
	parentPos := h.up(pos)
	for parentPos != -1 {
		if h.Tree[parentPos] <= h.Tree[pos] {
			break
		}
		swap := h.Tree[pos]
		h.Tree[pos] = h.Tree[parentPos]
		h.Tree[parentPos] = swap
		parentPos = h.up(parentPos)
	}
	fmt.Println(h.Tree)
}

func (h *MinHeap) heapifyDown(pos int) {
	left := h.downLeft(pos)
	right := h.downRight(pos)
	min := h.Tree[pos]
	idx := pos
	if left != -1 && min > h.Tree[left] {
		min = h.Tree[left]
		idx = left
	}
	if right != -1 && min > h.Tree[right] {
		min = h.Tree[right]
		idx = right
	}
	if min == h.Tree[pos] {
		return
	}
	swap := h.Tree[pos]
	h.Tree[pos] = h.Tree[idx]
	h.Tree[idx] = swap
	h.heapifyDown(idx)
}

// Insert inserts value to MinHeap
func (h *MinHeap) Insert(val int) {
	h.Tree = append(h.Tree, val)
	h.heapifyUp(h.len() - 1)
}

// Pop returns minimum value and remove it from heap
func (h *MinHeap) Pop() (int, error) {
	if h.len() <= 0 {
		return 0, fmt.Errorf("No value")
	}
	res := h.Tree[0]
	h.Tree[0] = h.Tree[h.len()-1]
	h.Tree = h.Tree[0 : h.len()-1]
	h.heapifyDown(0)
	return res, nil
}
