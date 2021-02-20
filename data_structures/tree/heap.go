package tree

import (
	"fmt"
)

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
	Insert(item Interface) // O(log n)
	Remove() Interface     // O(log n)
	Sort()                 // O(n*log n)
	Items() []Interface
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

// can only remove root element
func (h *heap) Remove() Interface {
	if len(h.items) == 0 {
		return nil
	}
	root := h.items[0]
	length := len(h.items)
	// swap root element with last element
	h.swap(0, length-1)
	// remove last element
	h.items = h.items[:length-1]
	h.heapifyDown(0, len(h.items))
	return root
}

// recursively swap root with largest child
// O(log n)
func (h *heap) heapifyDown(start, end int) {
	if start >= end || start < 0 || end == 0 {
		return
	}
	l, r := start*2+1, start*2+2
	max := start
	if l < end && h.items[max].Compare(h.items[l]) < 0 {
		max = l
	}
	if r < end && h.items[max].Compare(h.items[r]) < 0 {
		max = r
	}
	if max != start {
		h.swap(max, start)
		h.heapifyDown(max, end)
	}
}

func (h *heap) Sort() {
	lastIndex := len(h.items) - 1
	// O(n)
	for lastIndex > 0 {
		h.swap(0, lastIndex)
		lastIndex--
		// O(log n)
		h.heapifyDown(0, lastIndex)
	}
}

func (h *heap) Items() []Interface {
	return h.items
}

func (h *heap) String() string {
	return fmt.Sprint(h.items)
}
