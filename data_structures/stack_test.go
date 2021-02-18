package data_structures

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	q := NewStack()
	q.Push(1)
	q.Push(2)
	fmt.Println(q)
}

func TestPop(t *testing.T) {
	q := NewStack()
	q.Push(10)
	q.Push(20)
	q.Push(30)
	q.Push(40)
	var item int
	item, ok := q.Pop()
	fmt.Println(q, item, ok)
	item, ok = q.Pop()
	fmt.Println(q, item, ok)
	item, ok = q.Pop()
	fmt.Println(q, item, ok)
	item, ok = q.Pop()
	fmt.Println(q, item, ok)
}
