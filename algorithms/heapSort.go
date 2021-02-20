package algorithms

import "github.com/tao-yi/data-structure-in-go/data_structures/tree"

func HeapSort(arr []tree.Interface) []tree.Interface {
	h := tree.NewHeap()
	for _, item := range arr {
		h.Insert(item)
	}
	h.Sort()
	return h.Items()
}
