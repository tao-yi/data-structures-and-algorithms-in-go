package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		l     []int
		r     []int
		merge []int
	}{
		{l: []int{1, 3, 5, 7, 9}, r: []int{2, 3, 4, 6, 8}, merge: []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9}},
		{l: []int{1}, r: []int{2, 3}, merge: []int{1, 2, 3}},
		{l: []int{3, 5}, r: []int{2, 3}, merge: []int{2, 3, 3, 5}},
		{l: []int{}, r: []int{2}, merge: []int{2}},
		{l: []int{1}, r: []int{}, merge: []int{1}},
		{l: []int{}, r: []int{}, merge: []int{}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("merge %v, %v", test.l, test.r), func(t *testing.T) {
			actual := merge(test.l, test.r)
			if !reflect.DeepEqual(actual, test.merge) {
				t.Errorf("got %v, want %v", actual, test.merge)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	input := []int{9, 3, 7, 5, 6, 4, 8, 2}
	sorted := MergeSort(input)
	fmt.Println(sorted)
}
