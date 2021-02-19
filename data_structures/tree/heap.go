package tree

import "fmt"

/*
				A
		B				C
	D		E		F		G
			|0|1|2|3|4|5|6|
arr: [A,B,C,D,E,F,G]

left child: 	n*2 + 1
right child:  n*2 + 2
parent node:  (n-1) / 2
*/

type Heap interface {
	Insert(item Interface)
	Remove(item Interface) bool
}

type heap struct {
	items []Interface
}

func NewHeap() Heap {
	return &heap{}
}

func (h *heap) Insert(item Interface) {
	h.items = append(h.items, item)
	h.heapifyUp(len(h.items) - 1)
}

func (h *heap) heapifyUp(index int) {
	if index == 0 {
		return
	}

	parentIndex := (index - 1) / 2
	if h.items[index].Compare(h.items[parentIndex]) > 0 {
		h.swap(index, parentIndex)
		h.heapifyUp(parentIndex)
	}
}

func (h *heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *heap) Remove(item Interface) bool {
	return false
}

func (h *heap) String() string {
	return fmt.Sprint(h.items)
}
