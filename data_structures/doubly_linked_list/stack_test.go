package doubly_linked_list_test

import (
	"fmt"
	"testing"

	"github.com/tao-yi/data-structure-in-go/data_structures/doubly_linked_list"
)

func InitStack() doubly_linked_list.Stack {
	q := doubly_linked_list.NewStack()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	return q
}

func TestPush(t *testing.T) {
	s := doubly_linked_list.NewStack()
	tests := []struct {
		stack doubly_linked_list.Stack
		data  string
		size  int
	}{
		{stack: s, data: "hello", size: 1},
		{stack: s, data: "world", size: 2},
		{stack: s, data: "你好", size: 3},
		{stack: s, data: "世界", size: 4},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("enqueue %s", test.data), func(t *testing.T) {
			test.stack.Push(test.data)
			if test.stack.Size() != test.size {
				t.Errorf("got size: %d, want %d", test.stack.Size(), test.size)
			}
		})
	}
}

func TestPop(t *testing.T) {
	s := InitStack()
	tests := []struct {
		stack   doubly_linked_list.Stack
		data    interface{}
		size    int
		ok      bool
		isEmpty bool
	}{
		{stack: s, data: 4, size: 3, ok: true, isEmpty: false},
		{stack: s, data: 3, size: 2, ok: true, isEmpty: false},
		{stack: s, data: 2, size: 1, ok: true, isEmpty: false},
		{stack: s, data: 1, size: 0, ok: true, isEmpty: true},
		{stack: s, data: nil, size: 0, ok: false, isEmpty: true},
		{stack: s, data: nil, size: 0, ok: false, isEmpty: true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("pop from stack: %v", test.stack), func(t *testing.T) {
			data, ok := s.Pop()
			if test.stack.Size() != test.size {
				t.Errorf("got size: %d, want %d", test.stack.Size(), test.size)
			}
			if ok != test.ok {
				t.Errorf("got ok: %t, want %t", ok, test.ok)
			}
			if data != test.data {
				t.Errorf("got data: %v, want %v", data, test.data)
			}
			if test.stack.IsEmpty() != test.isEmpty {
				t.Errorf("got isEmpty: %t, want %t", test.stack.IsEmpty(), test.isEmpty)
			}
		})
	}
}
