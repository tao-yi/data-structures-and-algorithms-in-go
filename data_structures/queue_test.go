package data_structures

import (
	"fmt"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q)
}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	var item int
	item, ok := q.Dequeue()
	fmt.Println(q, item, ok)
	item, ok = q.Dequeue()
	fmt.Println(q, item, ok)
	item, ok = q.Dequeue()
	fmt.Println(q, item, ok)
	item, ok = q.Dequeue()
	fmt.Println(q, item, ok)
}
