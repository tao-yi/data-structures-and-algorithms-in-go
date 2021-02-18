package doubly_linked_list_test

import (
	"fmt"
	"testing"

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
		if item, ok := stack.Pop(); ok != false {
			arr[i] = item
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

// input: [[()]{()}]
func CheckForParentheses(input string) bool {
	opening := map[string]string{
		"[": "]",
		"(": ")",
		"{": "}",
	}
	closing := map[string]string{
		"]": "[",
		")": "(",
		"}": "{",
	}
	stack := []string{}
	var char string
	for _, c := range input {
		char = string(c)
		_, isOpening := opening[char]
		_, isClosing := closing[char]
		if isOpening {
			stack = append(stack, char)
		} else if isClosing {
			// pop last item from stack
			top := stack[len(stack)-1]
			if top == closing[char] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return true
}

func TestStack(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr1 = Reverse(arr1)
	arr2 = ReverseInPlace(arr2)
	fmt.Println(arr1)
	fmt.Println(arr2)

	input := "[[()]{()}]"
	isBalanced := CheckForParentheses(input)
	fmt.Println(isBalanced)
	input = "[[((((({)))))]{()}]"
	isBalanced = CheckForParentheses(input)
	fmt.Println(isBalanced)
}
