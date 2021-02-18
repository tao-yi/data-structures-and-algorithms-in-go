package data_structures

import "testing"

func TestAdd(t *testing.T) {
	// tests := []struct {
	// 	input Bag
	// 	item  int
	// 	want  Bag
	// }{
	// 	{&bag{[]int{1, 2, 3}}, 4, &bag{[]int{1, 2, 3, 4}}},
	// 	{&bag{[]int{1, 2}}, 3, &bag{[]int{1, 2, 3}}},
	// 	{&bag{[]int{1}}, 2, &bag{[]int{1, 2}}},
	// 	{&bag{[]int{}}, 1, &bag{[]int{1}}},
	// }

	// for _, test := range tests {
	// 	test.input.Add(test.item)
	// 	if test.input != test.want {
	// 		t.Errorf("got %v, want %v", test.input, test.want)
	// 	}
	// }
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input Bag
		want  bool
	}{
		{&bag{[]int{1, 2, 3}}, false},
		{&bag{[]int{1, 2}}, false},
		{&bag{[]int{}}, true},
	}

	for _, test := range tests {
		actual := test.input.IsEmpty()
		if actual != test.want {
			t.Errorf("%v got %t, want %t", test.input, actual, test.want)
		}
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		input Bag
		want  int
	}{
		{&bag{[]int{1, 2, 3}}, 3},
		{&bag{[]int{1, 2}}, 2},
		{&bag{[]int{1}}, 1},
	}

	for _, test := range tests {
		actual := test.input.Size()
		if actual != test.want {
			t.Errorf("%v got %d, want %d", test.input, actual, test.want)
		}
	}
}
