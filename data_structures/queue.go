package data_structures

type Queue interface {
	Enqueue(interface{})          // O(1)
	Dequeue() (interface{}, bool) // O(1)
	IsEmpty() bool                // O(1)
	Size() int                    // O(1)
}

type queue struct {
	items []interface{}
}

func NewQueue() Queue {
	return &queue{}
}

func (q *queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *queue) Dequeue() (interface{}, bool) {
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
