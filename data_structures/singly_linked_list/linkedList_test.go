package singly_linked_list

import (
	"testing"
)

func TestAddFirst(t *testing.T) {
	l := NewLinkedList()
	l.AddFirst(1)
	Debug(l)
	l.AddFirst(2)
	Debug(l)
	l.AddFirst(3)
	Debug(l)
}

func TestAddLast(t *testing.T) {
	l := NewLinkedList()
	l.AddLast(1)
	Debug(l)
	l.AddLast(2)
	Debug(l)
	l.AddLast(3)
	Debug(l)
}

func TestRemoveFirst(t *testing.T) {
	l := NewLinkedList()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	var tests = []struct {
		input   LinkedList
		item    int
		success bool
	}{
		{l, 1, true},
		{l, 2, true},
		{l, 3, true},
		{l, 0, false},
	}

	for _, test := range tests {
		item, ok := l.RemoveFirst()
		Debug(l)
		if item != test.item {
			t.Errorf("got %d, want %d", item, test.item)
		}
		if ok != test.success {
			t.Errorf("got %t, want %t", ok, test.success)
		}
	}
}

func TestRemoveLast(t *testing.T) {
	l := NewLinkedList()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)

	var tests = []struct {
		input   LinkedList
		item    int
		success bool
	}{
		{l, 3, true},
		{l, 2, true},
		{l, 1, true},
		{l, 0, false},
	}

	for _, test := range tests {
		item, ok := l.RemoveLast()
		Debug(l)
		if item != test.item {
			t.Errorf("got %d, want %d", item, test.item)
		}
		if ok != test.success {
			t.Errorf("got %t, want %t", ok, test.success)
		}
	}
}

func TestRemove(t *testing.T) {
	l := NewLinkedList()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	l.AddLast(4)
	Debug(l)

	var tests = []struct {
		input   LinkedList
		item    int
		success bool
	}{
		{l, 3, true},
		{l, 2, true},
		{l, 1, true},
		{l, 0, false},
		{l, 4, true},
	}

	for _, test := range tests {
		item, ok := l.Remove(test.item)
		Debug(l)
		if item != test.item {
			t.Errorf("got %d, want %d", item, test.item)
		}
		if ok != test.success {
			t.Errorf("got %t, want %t", ok, test.success)
		}
	}

	Debug(l)
}

func TestContains(t *testing.T) {
	l := NewLinkedList()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	l.AddLast(4)

	var tests = []struct {
		input LinkedList
		item  int
		found bool
	}{
		{l, 3, true},
		{l, 2, true},
		{l, 1, true},
		{l, 0, false},
		{l, 6, false},
		{l, 4, true},
	}

	for _, test := range tests {
		found := l.Contains(test.item)
		if found != test.found {
			t.Errorf("got %t, want %t", found, test.found)
		}
	}
}
