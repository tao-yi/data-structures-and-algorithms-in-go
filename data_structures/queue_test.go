package data_structures

import (
	"fmt"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := NewQueue()

	tests := []struct {
		q       Queue
		item    int
		size    int
		isEmpty bool
	}{
		{q, 1, 1, false},
		{q, 2, 2, false},
		{q, 3, 3, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("enqueue %d", test.item), func(t *testing.T) {
			q.Enqueue(test.item)
			size := q.Size()
			isEmpty := q.IsEmpty()
			if size != test.size {
				t.Errorf("got size: %d, want %d", size, test.size)
			}

			if isEmpty != test.isEmpty {
				t.Errorf("got isEmpty: %t, want %t", isEmpty, test.isEmpty)
			}
		})
	}

}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	var item interface{}
	item, ok := q.Dequeue()
	fmt.Println(q, item, ok)
	item, ok = q.Dequeue()
	fmt.Println(q, item, ok)
	item, ok = q.Dequeue()
	fmt.Println(q, item, ok)
	item, ok = q.Dequeue()
	fmt.Println(q, item, ok)
}
