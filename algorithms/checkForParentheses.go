package algorithms

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
