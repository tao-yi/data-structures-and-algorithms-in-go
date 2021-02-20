package dp

import (
	"fmt"
	"testing"
)

func TestFib1(t *testing.T) {
	tests := []struct {
		n      int
		result int
	}{
		{n: 0, result: 0},
		{n: 1, result: 1},
		{n: 2, result: 1},
		{n: 3, result: 2},
		{n: 4, result: 3},
		{n: 5, result: 5},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("f(%d)", test.n), func(t *testing.T) {
			res := Fib1(test.n)
			if res != test.result {
				t.Errorf("got %d, want %d", res, test.result)
			}
		})
	}
}

func TestFib2(t *testing.T) {
	tests := []struct {
		n      int
		result int
	}{
		{n: 0, result: 0},
		{n: 1, result: 1},
		{n: 2, result: 1},
		{n: 3, result: 2},
		{n: 4, result: 3},
		{n: 5, result: 5},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("f(%d)", test.n), func(t *testing.T) {
			res := Fib1(test.n)
			if res != test.result {
				t.Errorf("got %d, want %d", res, test.result)
			}
		})
	}
}
