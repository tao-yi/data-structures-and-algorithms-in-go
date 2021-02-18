package doubly_linked_list

type Stack interface {
	Push(int)         // O(1)
	Pop() (int, bool) // O(1)
	Top() (int, bool) // O(1)
	IsEmpty() bool    // O(1)
	Size() int        // O(1)
}

type stack struct {
	items DoublyLinkedlist
}

func NewStack() Stack {
	return &stack{
		items: &doublyLinkedList{},
	}
}

func (s *stack) Push(item int) {
	s.items.AddLast(item)
}

func (s *stack) Pop() (int, bool) {
	return s.items.RemoveLast()
}

func (s *stack) Top() (int, bool) {
	return s.items.PeekLast()
}

func (s *stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stack) Size() int {
	return s.items.Size()
}
