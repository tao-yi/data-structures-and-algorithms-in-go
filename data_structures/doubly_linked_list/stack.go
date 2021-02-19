package doubly_linked_list

type Stack interface {
	Push(interface{})         // O(1)
	Pop() (interface{}, bool) // O(1)
	Top() (interface{}, bool) // O(1)
	IsEmpty() bool            // O(1)
	Size() int                // O(1)
}

type stack struct {
	items DoublyLinkedlist
}

func NewStack() Stack {
	return &stack{
		items: &doublyLinkedList{},
	}
}

func (s *stack) Push(item interface{}) {
	s.items.AddLast(item)
}

func (s *stack) Pop() (interface{}, bool) {
	return s.items.RemoveLast()
}

func (s *stack) Top() (interface{}, bool) {
	return s.items.PeekLast()
}

func (s *stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stack) Size() int {
	return s.items.Size()
}
