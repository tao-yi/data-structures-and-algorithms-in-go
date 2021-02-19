package data_structures

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		input Bag
		item  int
	}{
		{&bag{}, 4},
		{&bag{}, 3},
		{&bag{}, 2},
		{&bag{}, 1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Add: %d", test.item), func(t *testing.T) {
			test.input.Add(test.item)
			var include bool
			for _, item := range test.input.Values() {
				if item.(int) == test.item {
					include = true
				}
			}
			if !include {
				t.Errorf("added item %d should be in bag %v", test.item, test.input)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input Bag
		want  bool
	}{
		{&bag{[]interface{}{1, 2, 3}}, false},
		{&bag{[]interface{}{1, 2}}, false},
		{&bag{[]interface{}{}}, true},
	}

	for _, test := range tests {
		actual := test.input.IsEmpty()
		t.Run(fmt.Sprintf("test %v isEmpty", test.input), func(t *testing.T) {
			if actual != test.want {
				t.Errorf("%v got %t, want %t", test.input, actual, test.want)
			}
		})
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		input Bag
		want  int
	}{
		{&bag{[]interface{}{1, 2, 3}}, 3},
		{&bag{[]interface{}{1, 2}}, 2},
		{&bag{[]interface{}{1}}, 1},
	}

	for _, test := range tests {
		actual := test.input.Size()
		t.Run(fmt.Sprintf("size of %v", test.input), func(t *testing.T) {
			if actual != test.want {
				t.Errorf("%v got %d, want %d", test.input, actual, test.want)
			}
		})
	}
}
