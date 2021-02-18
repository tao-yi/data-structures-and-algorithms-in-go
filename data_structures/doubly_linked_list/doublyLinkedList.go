package doubly_linked_list

import "fmt"

type node struct {
	data int
	prev *node
	next *node
}

type DoublyLinkedlist interface {
	AddFirst(data int)           // O(1)
	AddLast(data int)            // O(1)
	RemoveFirst() (int, bool)    // O(1)
	RemoveLast() (int, bool)     // O(1)
	Remove(item int) (int, bool) // O(n)
	Contains(item int) bool      // O(n)
	PeekFirst() (int, bool)      // O(1)
	PeekLast() (int, bool)       // O(1)
	Size() int                   // O(1)
}

func newNode(data int) *node {
	return &node{data: data, prev: nil, next: nil}
}

type doublyLinkedList struct {
	head *node
	tail *node
	size int
}

func (l *doublyLinkedList) AddFirst(data int) {
	newNode := newNode(data)
	newNode.next = l.head
	if l.head == nil {
		l.tail = newNode
	}
	l.head.prev = newNode
	l.head = newNode
	l.size++
}

func (l *doublyLinkedList) AddLast(data int) {
	newNode := newNode(data)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
	l.size++
}

func (l *doublyLinkedList) RemoveFirst() (first int, success bool) {
	// empty list
	if l.head == nil {
		return
	}
	first = l.head.data
	// a single element list
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
	}

	l.size--
	success = true
	return first, success
}

func (l *doublyLinkedList) RemoveLast() (last int, success bool) {
	if l.head == nil {
		return
	}
	if l.head == l.tail {
		return l.RemoveFirst()
	}
	last = l.tail.data
	l.tail.prev.next = nil
	l.tail = l.tail.prev
	l.size--
	success = true
	return last, success
}

func (l *doublyLinkedList) Remove(item int) (int, bool) {
	var previousNode *node = nil
	var currentNode *node = l.head
	for currentNode != nil {
		if currentNode.data == item {
			if currentNode == l.head {
				return l.RemoveFirst()
			}
			if currentNode == l.tail {
				return l.RemoveLast()
			}
			// delete the node
			previousNode.next = currentNode.next
			currentNode.next.prev = previousNode
			l.size--
			return currentNode.data, true
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}

	return 0, false
}

func (l *doublyLinkedList) Contains(item int) bool {
	var currentNode *node = l.head
	for currentNode != nil {
		if currentNode.data == item {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (l *doublyLinkedList) PeekFirst() (int, bool) {
	if l.head == nil {
		return 0, false
	}
	return l.head.data, true
}

func (l *doublyLinkedList) PeekLast() (int, bool) {
	if l.tail == nil {
		return 0, false
	}
	return l.tail.data, true
}

func (l *doublyLinkedList) Size() int {
	return l.size
}

func Debug(l DoublyLinkedlist) {
	head := l.(*doublyLinkedList).head
	if head == nil {
		fmt.Println(l.Size())
		return
	}
	currentNode := l.(*doublyLinkedList).head
	for currentNode != nil {
		fmt.Printf("[data:%d]", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println(l.Size())
}
