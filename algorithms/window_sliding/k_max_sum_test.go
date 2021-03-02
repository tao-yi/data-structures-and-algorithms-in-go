package window_sliding

import (
	"testing"
)

/*
Given an array of integers of size ‘n’.
Our aim is to calculate the maximum sum of ‘k’
consecutive elements in the array.
*/

// O(n*k)
func kMaxBruteForce(input []int, k int) int {
	var max int
	for i, j := 0, k-1; j < len(input); i, j = i+1, j+1 {
		var sum int
		for l := i; l <= j; l++ {
			sum += input[l]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// O(n)
func kMaxWindowSliding(input []int, k int) int {
	var max, sum int
	// create the first window sum
	for i := 0; i < k; i++ {
		sum += input[i]
	}
	max = sum
	// now we have a window [0, k-1]
	// sliding the window from position k to the right
	// one position at a time
	for i := k; i < len(input); i++ {
		// Compute sums of remaining windows by
		// removing first element of previous
		// window and adding last element of
		// current window.
		sum = sum + input[i] - input[i-k]
		if sum > max {
			max = sum
		}
	}
	return max
}

func TestKMax(t *testing.T) {
	tests := []struct {
		input []int
		k     int
		sum   int
	}{
		{input: []int{100, 200, 300, 400}, k: 2, sum: 700},
		{input: []int{1, 4, 2, 10, 23, 3, 1, 0, 20}, k: 4, sum: 39},
	}

	for _, test := range tests {
		output := kMaxBruteForce(test.input, test.k)
		if test.sum != output {
			t.Errorf("got %d, want %d", output, test.sum)
		}

		output2 := kMaxWindowSliding(test.input, test.k)
		if test.sum != output2 {
			t.Errorf("got %d, want %d", output2, test.sum)
		}
	}
}
