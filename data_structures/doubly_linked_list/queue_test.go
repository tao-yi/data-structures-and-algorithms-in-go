package doubly_linked_list_test

import (
	"fmt"
	"testing"

	"github.com/tao-yi/data-structure-in-go/data_structures/doubly_linked_list"
)

func InitQueue() doubly_linked_list.Queue {
	q := doubly_linked_list.NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	return q
}

func TestEnqueue(t *testing.T) {
	q := doubly_linked_list.NewQueue()

	tests := []struct {
		queue doubly_linked_list.Queue
		data  string
		size  int
	}{
		{queue: q, data: "hello", size: 1},
		{queue: q, data: "world", size: 2},
		{queue: q, data: "你好", size: 3},
		{queue: q, data: "世界", size: 4},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("enqueue %s", test.data), func(t *testing.T) {
			test.queue.Enqueue(test.data)
			if test.queue.Size() != test.size {
				t.Errorf("got size: %d, want %d", test.queue.Size(), test.size)
			}
		})
	}
}

func TestDequeue(t *testing.T) {
	q := InitQueue()

	tests := []struct {
		queue   doubly_linked_list.Queue
		data    interface{}
		size    int
		ok      bool
		isEmpty bool
	}{
		{queue: q, data: 1, size: 3, ok: true, isEmpty: false},
		{queue: q, data: 2, size: 2, ok: true, isEmpty: false},
		{queue: q, data: 3, size: 1, ok: true, isEmpty: false},
		{queue: q, data: 4, size: 0, ok: true, isEmpty: true},
		{queue: q, data: nil, size: 0, ok: false, isEmpty: true},
		{queue: q, data: nil, size: 0, ok: false, isEmpty: true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("dequeue %v", test.queue), func(t *testing.T) {
			data, ok := test.queue.Dequeue()
			if test.queue.Size() != test.size {
				t.Errorf("got size: %d, want %d", test.queue.Size(), test.size)
			}
			if ok != test.ok {
				t.Errorf("got ok: %t, want %t", ok, test.ok)
			}
			if data != test.data {
				t.Errorf("got data: %v, want %v", data, test.data)
			}
			if test.queue.IsEmpty() != test.isEmpty {
				t.Errorf("got isEmpty: %t, want %t", q.IsEmpty(), test.isEmpty)
			}
		})
	}
}
