package data_structures

type Queue interface {
	Enqueue(int)          // O(1)
	Dequeue() (int, bool) // O(1)
	IsEmpty() bool        // O(1)
	Size() int            // O(1)
}

type queue struct {
	items []int
}

func NewQueue() Queue {
	return &queue{}
}

func (q *queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	head := q.items[0]
	q.items = q.items[1:]
	return head, true
}

func (q *queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *queue) Size() int {
	return len(q.items)
}
