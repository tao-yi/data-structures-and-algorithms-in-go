package tree

import (
	"fmt"
	"testing"

	"github.com/tao-yi/data-structure-in-go/util"
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

func InitHeap() Heap {
	h := NewHeap()
	h.Insert(NewHeapItem(1))
	h.Insert(NewHeapItem(2))
	h.Insert(NewHeapItem(3))
	h.Insert(NewHeapItem(4))
	h.Insert(NewHeapItem(5))
	h.Insert(NewHeapItem(6))
	return h
}

func TestRemove(t *testing.T) {
	h := InitHeap()
	tests := []struct {
		heap Heap
		item Interface
	}{
		{heap: h, item: NewHeapItem(6)},
		{heap: h, item: NewHeapItem(5)},
		{heap: h, item: NewHeapItem(4)},
		{heap: h, item: NewHeapItem(3)},
		{heap: h, item: NewHeapItem(2)},
		{heap: h, item: NewHeapItem(1)},
		{heap: h, item: nil},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("remove %v from heap", test.item), func(t *testing.T) {
			item := test.heap.Remove()
			if item != test.item {
				t.Errorf("got %v, want %v", item, test.item)
			}
		})
	}
}

func TestSort(t *testing.T) {
	h := InitHeap()
	util.MakeImg("heap", &h)
	fmt.Println(h)
	h.Sort()
	util.MakeImg("heap-sorted", &h)
	fmt.Println(h)
}
