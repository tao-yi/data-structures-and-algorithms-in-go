package doubly_linked_list

import (
	"fmt"
	"testing"
)

func TestRemoveFirst(t *testing.T) {
	list := NewDoublyLinkedlist()
	list.AddFirst("hello")

	// tests := []struct {
	// 	input DoublyLinkedlist
	// 	item  string
	// 	ok    bool
	// }{
	// 	{list, "hello", true},
	// 	{list, "", false},
	// 	{list, "", false},
	// }

	// first, ok := list.RemoveFirst()
	// if first != "hello" {
	// 	t.Errorf("got %s, %t want %s, %t", first, ok, test.item, test.ok)
	// }

	item, ok := list.RemoveFirst()
	fmt.Println(item, ok)
	item, ok = list.RemoveFirst()
	fmt.Println(item, ok)
	item, ok = list.RemoveLast()
	fmt.Println(item, ok)
}

func TestRemoveLast(t *testing.T) {
	list := NewDoublyLinkedlist()
	list.AddFirst("hello")
	list.AddFirst("world")
	list.AddFirst("你好")

	item, ok := list.RemoveLast()
	fmt.Println(item, ok)
	item, ok = list.RemoveLast()
	fmt.Println(item, ok)
	item, ok = list.RemoveLast()
	fmt.Println(item, ok)
	item, ok = list.RemoveLast()
	fmt.Println(item, ok)
}
