package tree

import (
	"fmt"
	"testing"
)

type heapItem int

func NewHeapItem(data int) heapItem {
	return heapItem(data)
}
func (h heapItem) Compare(data interface{}) int {
	return int(h) - int(data.(heapItem))
}

func TestInsert(t *testing.T) {
	h := NewHeap()

	h.Insert(NewHeapItem(1))
	h.Insert(NewHeapItem(2))
	h.Insert(NewHeapItem(3))
	h.Insert(NewHeapItem(4))
	h.Insert(NewHeapItem(5))
	h.Insert(NewHeapItem(6))
	fmt.Println(h)
}
