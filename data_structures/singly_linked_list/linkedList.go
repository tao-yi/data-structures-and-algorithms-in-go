package singly_linked_list

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList interface {
	AddFirst(data int)           // O(1)
	AddLast(data int)            // O(1)
	RemoveFirst() (int, bool)    // O(1)
	RemoveLast() (int, bool)     // O(n)
	Remove(item int) (int, bool) // O(n)
	Contains(item int) bool      // O(n)
	PeekFirst() (int, bool)      // O(1)
	PeekLast() (int, bool)       // O(1)
	Size() int                   // O(1)
}

func newNode(data int) *node {
	return &node{data: data, next: nil}
}

type linkedList struct {
	head *node
	tail *node
	size int
}

func NewLinkedList() LinkedList {
	return &linkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l *linkedList) AddFirst(data int) {
	newNode := newNode(data)
	newNode.next = l.head
	if l.head == nil {
		l.tail = newNode
	}
	l.head = newNode
	l.size++
}

func (l *linkedList) AddLast(data int) {
	newNode := newNode(data)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

func (l *linkedList) RemoveFirst() (first int, success bool) {
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

func (l *linkedList) RemoveLast() (last int, success bool) {
	if l.head == nil {
		return
	}
	if l.head == l.tail {
		return l.RemoveFirst()
	}
	var previous *node = nil
	var current *node = l.head
	for current != l.tail {
		previous = current
		current = current.next
	}
	previous.next = nil
	l.tail = previous
	l.size--

	last = current.data
	success = true
	return last, success
}

func (l *linkedList) Remove(item int) (int, bool) {
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
			l.size--
			return currentNode.data, true
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}

	return 0, false
}

func (l *linkedList) Contains(item int) bool {
	var currentNode *node = l.head
	for currentNode != nil {
		if currentNode.data == item {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (l *linkedList) PeekFirst() (int, bool) {
	if l.head == nil {
		return 0, false
	}
	return l.head.data, true
}

func (l *linkedList) PeekLast() (int, bool) {
	if l.tail == nil {
		return 0, false
	}
	return l.tail.data, true
}

func (l *linkedList) Size() int {
	return l.size
}

func Debug(l LinkedList) {
	head := l.(*linkedList).head
	if head == nil {
		fmt.Println(l.Size())
		return
	}
	currentNode := l.(*linkedList).head
	for currentNode != nil {
		fmt.Printf("[data:%d]", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println(l.Size())
}
