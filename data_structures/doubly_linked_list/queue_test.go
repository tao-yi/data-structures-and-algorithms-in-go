package doubly_linked_list

import (
	"fmt"
	"testing"
)

func TestDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item, ok := q.Dequeue()
	fmt.Println(item, ok)
	item, ok = q.Dequeue()
	fmt.Println(item, ok)
	item, ok = q.Dequeue()
	fmt.Println(item, ok)
	item, ok = q.Dequeue()
	fmt.Println(item, ok)
}
