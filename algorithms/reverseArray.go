package algorithms

import (
	"fmt"

	"github.com/tao-yi/data-structure-in-go/data_structures/doubly_linked_list"
)

// using stack to reverse
func Reverse(arr []int) []int {
	stack := doubly_linked_list.NewStack()
	for _, item := range arr {
		stack.Push(item)
	}
	i := 0
	for !stack.IsEmpty() {
		if item, ok := stack.Pop(); !ok {
			arr[i] = item.(int)
			i++
		}
	}
	fmt.Println(stack.Size())
	return arr
}

// reverse with space complexity O(1)
func ReverseInPlace(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
