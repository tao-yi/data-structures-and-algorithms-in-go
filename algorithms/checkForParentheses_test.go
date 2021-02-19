package algorithms

import (
	"fmt"
	"testing"
)

func TestCheckForParentheses(t *testing.T) {
	tests := []struct {
		input  string
		result bool
	}{
		{input: "[a[b()]{()}]", result: true},
		{input: "[[[(c)]{()}]", result: false},
		{input: "[[[(e)]{(f)}]]]", result: false},
		{input: "[[[()]{()}g]]", result: true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("validate %s", test.input), func(t *testing.T) {
			res := CheckForParentheses(test.input)
			if res != test.result {
				t.Errorf("got %t, want %t", res, test.result)
			}
		})
	}
}
