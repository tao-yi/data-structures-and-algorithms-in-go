package doubly_linked_list

type Queue interface {
	Enqueue(interface{})          // O(1)
	Dequeue() (interface{}, bool) // O(1)
	IsEmpty() bool                // O(1)
	Size() int                    // O(1)
}

type queue struct {
	items DoublyLinkedlist
}

func NewQueue() Queue {
	return &queue{
		items: &doublyLinkedList{},
	}
}

func (q *queue) Enqueue(item interface{}) {
	q.items.AddLast(item)
}

func (q *queue) Dequeue() (interface{}, bool) {
	return q.items.RemoveFirst()
}

func (q *queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *queue) Size() int {
	return q.items.Size()
}
