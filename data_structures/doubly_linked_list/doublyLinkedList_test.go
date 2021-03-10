package doubly_linked_list_test

import (
	"fmt"
	"testing"

	"github.com/tao-yi/data-structure-in-go/data_structures/doubly_linked_list"
	"github.com/tao-yi/data-structure-in-go/util"
)

func TestAddFirst(t *testing.T) {
	list := doubly_linked_list.NewDoublyLinkedlist()
	tests := []struct {
		list doubly_linked_list.DoublyLinkedlist
		data string
		size int
	}{
		{list, "hello", 1},
		{list, "world", 2},
		{list, "你好", 3},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("AddFirst %s", test.data), func(t *testing.T) {
			list.AddFirst(test.data)
			if list.Size() != test.size {
				t.Errorf("got size %d, want %d", list.Size(), test.size)
			}
			first, ok := list.PeekFirst()
			if first != test.data {
				t.Errorf("got first %s, want %s", first, test.data)
			}
			if !ok {
				t.Errorf("got ok %t, want true", ok)
			}
		})
	}
}

func InitList() doubly_linked_list.DoublyLinkedlist {
	list := doubly_linked_list.NewDoublyLinkedlist()
	list.AddFirst("hello")
	list.AddFirst("world")
	list.AddFirst("你好")
	return list
}

func TestRemoveFirst(t *testing.T) {
	list := InitList()
	util.MakeImg("doublyLinkedlist", &list)
	tests := []struct {
		input doubly_linked_list.DoublyLinkedlist
		item  interface{}
		ok    bool
		size  int
	}{
		{list, "你好", true, 2},
		{list, "world", true, 1},
		{list, "hello", true, 0},
		{list, nil, false, 0},
		{list, nil, false, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("removeFirst from %v", test.input), func(t *testing.T) {
			first, ok := list.RemoveFirst()
			if first != test.item {
				t.Errorf("got first %s want %s", first, test.item)
			}
			if ok != test.ok {
				t.Errorf("got ok %t, want %t", ok, test.ok)
			}
			if test.input.Size() != test.size {
				t.Errorf("got size %d, want %d", test.input.Size(), test.size)
			}
		})
	}
}

func TestRemoveLast(t *testing.T) {
	list := InitList()
	tests := []struct {
		input doubly_linked_list.DoublyLinkedlist
		item  interface{}
		ok    bool
		size  int
	}{
		{list, "hello", true, 2},
		{list, "world", true, 1},
		{list, "你好", true, 0},
		{list, nil, false, 0},
		{list, nil, false, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("removeFirst from %v", test.input), func(t *testing.T) {
			first, ok := list.RemoveLast()
			if first != test.item {
				t.Errorf("got first %s want %s", first, test.item)
			}
			if ok != test.ok {
				t.Errorf("got ok %t, want %t", ok, test.ok)
			}
		})
	}
}
