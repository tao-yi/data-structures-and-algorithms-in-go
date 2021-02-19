package algorithms

import (
	"fmt"

	"github.com/tao-yi/data-structure-in-go/data_structures/doubly_linked_list"
)

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
	stack := doubly_linked_list.NewStack()
	var char string
	for _, c := range input {
		char = string(c)
		_, isOpening := opening[char]
		_, isClosing := closing[char]
		if isOpening {
			stack.Push(char)
		} else if isClosing {
			// pop last item from stack
			top, ok := stack.Top()
			fmt.Printf("top:%v, char:%v, ok: %t\n", top, char, ok)
			if !ok {
				return false
			}
			if top == closing[char] {
				stack.Pop()
			} else {
				return false
			}
		} else {
			fmt.Printf("else char:%v\n", char)
		}
	}
	// if stack still has value means there are more opening then closing
	return stack.IsEmpty()
}
