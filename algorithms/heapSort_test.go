package algorithms

import (
	"fmt"
	"testing"

	"github.com/tao-yi/data-structure-in-go/data_structures/tree"
)

type heapData int

func newHeapData(data int) heapData {
	return heapData(data)
}

func (h heapData) Compare(data interface{}) int {
	return int(h) - int(data.(heapData))
}

func TestHeapSort(t *testing.T) {
	input := []int{1, 6, 16, 52, 3, 5, 9, 10, 25, 66, 33, 1}
	input2 := []tree.Interface{}
	for _, d := range input {
		input2 = append(input2, newHeapData(d))
	}

	input2 = HeapSort(input2)
	fmt.Println(input2)
}
