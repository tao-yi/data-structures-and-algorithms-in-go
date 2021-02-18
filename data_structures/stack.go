package data_structures

type Stack interface {
	Push(int)         // O(1)
	Pop() (int, bool) // O(1)
	Top() int         // O(1)
	IsEmpty() bool    // O(1)
	Size() int        // O(1)
}

type stack struct {
	items []int
}

func NewStack() Stack {
	return &stack{}
}

func (s *stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *stack) Top() int {
	return s.items[len(s.items)-1]
}

func (s *stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stack) Size() int {
	return len(s.items)
}
