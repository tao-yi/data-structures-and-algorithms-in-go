package doubly_linked_list

type Queue interface {
	Enqueue(int)          // O(1)
	Dequeue() (int, bool) // O(1)
	IsEmpty() bool        // O(1)
	Size() int            // O(1)
}
