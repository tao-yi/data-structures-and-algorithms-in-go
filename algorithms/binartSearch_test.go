package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		input  []int
		target int
		want   bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 6, true},
		{[]int{1, 3, 5, 7, 15, 62, 79}, 25, false},
		{[]int{1, 2, 5, 9, 21, 39, 50, 66, 71}, 2, true},
	}

	for _, test := range tests {
		output := BinarySearch(test.input, test.target)
		if test.want != output {
			t.Errorf("%v: got %t, want %t", test.input, output, test.want)
		}
	}
}
