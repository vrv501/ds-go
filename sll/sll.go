package sll

import "fmt"

type SinglyLinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length uint
	// For most primitive types(integers, decimal, string), cmp.Compare can be used
	orderedFunc func(a, b T) int // should return 0 if equal, -ve value if a < b, and +ve if a > b

}

type node[T any] struct {
	value T
	next  *node[T]
}

func NewSinglyLinkedList[T any](f func(a, b T) int) *SinglyLinkedList[T] {
	if f == nil {
		panic("Comparison function cannot be nil")
	}
	return &SinglyLinkedList[T]{orderedFunc: f}
}

func (sll *SinglyLinkedList[T]) Append(value T) {
	if sll == nil || sll.orderedFunc == nil {
		panic("Singly Linked List is not initialized")
	}

	sll.length++
	newNode := &node[T]{value: value}
	if sll.length == 1 {
		sll.head = newNode
		sll.tail = newNode
		return
	}

	sll.tail.next = newNode
	sll.tail = newNode
}

func (sll *SinglyLinkedList[T]) Prepend(value T) {
	if sll == nil || sll.orderedFunc == nil {
		panic("Singly Linked List is not initialized")
	}

	sll.length++
	newNode := &node[T]{value: value, next: sll.head}
	if sll.length == 1 {
		sll.head = newNode
		sll.tail = newNode
		return
	}

	sll.head = newNode
}

// Indexing starts at 0. If index > SLL length, function panics
func (sll *SinglyLinkedList[T]) InsertAtIndex(value T, index uint) {
	if sll == nil || sll.orderedFunc == nil {
		panic("Singly Linked List is not initialized")
	} else if index == 0 {
		sll.Prepend(value)
		return
	} else if index == sll.length {
		sll.Append(value)
		return
	} else if index > sll.length {
		panic("Index out of bounds")
	}

	curr := sll.head
	var prev *node[T]
	for index > 0 {
		prev = curr
		curr = curr.next
		index--
	}

	sll.length++
	newNode := &node[T]{value: value, next: curr}
	prev.next = newNode
}

func (sll *SinglyLinkedList[T]) Print() {
	if sll == nil || sll.orderedFunc == nil {
		panic("Singly Linked List is not initialized")
	} else if sll.length == 0 {
		return
	}

	fmt.Print(sll.head.value)
	curr := sll.head.next
	for curr != nil {
		fmt.Printf(" -> %v", curr.value)
		curr = curr.next
	}
	fmt.Println()
}

func (sll *SinglyLinkedList[T]) Length() uint {
	if sll == nil || sll.orderedFunc == nil {
		panic("Singly Linked List is not initialized")
	}

	return sll.length
}

func (sll *SinglyLinkedList[T]) Pop() T {
	if sll == nil || sll.orderedFunc == nil {
		panic("Singly Linked List is not initialized")
	} else if sll.length == 0 {
		panic("Pop from Empty Singly Linked List")
	}

	sll.length--
	val := sll.tail.value
	if sll.length == 0 {
		sll.head = nil
		sll.tail = nil
		return val
	}

	curr := sll.head
	var prev *node[T]
	for curr.next != nil {
		prev = curr
		curr = curr.next
	}

	sll.tail = prev
	sll.tail.next = nil
	return val
}

func (sll *SinglyLinkedList[T]) PopFromFront() T {
	if sll == nil || sll.orderedFunc == nil {
		panic("Singly Linked List is not initialized")
	} else if sll.length == 0 {
		panic("PopFromFront from Empty Singly Linked List")
	}

	sll.length--
	val := sll.head.value
	if sll.length == 0 {
		sll.head = nil
		sll.tail = nil
		return val
	}

	tmp := sll.head
	sll.head = sll.head.next
	tmp.next = nil
	return val
}

// Returns value at index after removing it. If index > length or sll length == 0, function panics
func (sll *SinglyLinkedList[T]) RemoveAtIndex(index uint) T {
	if sll == nil || sll.orderedFunc == nil {
		panic("Singly Linked List is not initialized")
	} else if sll.length == 0 {
		panic("RemoveAtIndex From Empty Singly Linked List")
	} else if index == 0 {
		return sll.PopFromFront()
	} else if index == sll.length-1 {
		return sll.Pop()
	} else if index >= sll.length {
		panic("Index out of bounds")
	}

	curr := sll.head
	var prev *node[T]
	sll.length--
	for index > 0 {
		prev = curr
		curr = curr.next
		index--
	}

	prev.next = curr.next
	curr.next = nil
	return curr.value
}

// Returns true when value is found and removed
func (sll *SinglyLinkedList[T]) RemoveVal(val T) bool {
	if sll == nil || sll.orderedFunc == nil {
		panic("Singly Linked List is not initialized")
	} else if sll.length == 0 {
		return false
	}

	curr := sll.head
	var prev *node[T]
	for curr != nil && sll.orderedFunc(val, curr.value) != 0 {
		prev = curr
		curr = curr.next
	}

	switch {
	case curr == nil:
		return false
	case curr == sll.head: // comparing address stored in pointers
		_ = sll.PopFromFront()
		return true
	case curr.next == nil:
		_ = sll.Pop()
		return true
	}

	sll.length--
	prev.next = curr.next
	curr.next = nil
	return true
}

// Returns index if value is found
func (sll *SinglyLinkedList[T]) Contains(val T) (uint, bool) {
	if sll == nil || sll.orderedFunc == nil {
		panic("Singly Linked List is not initialized")
	} else if sll.length == 0 {
		return 0, false
	}

	curr := sll.head
	var index uint
	for curr != nil && sll.orderedFunc(val, curr.value) != 0 {
		curr = curr.next
		index++
	}

	if curr == nil {
		return 0, false
	}
	return index, true
}

func (sll *SinglyLinkedList[T]) Reverse() {
	if sll == nil || sll.orderedFunc == nil {
		panic("Singly Linked List is not initialized")
	} else if sll.length <= 1 {
		return
	}

	curr := sll.head
	var next, prev *node[T]
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	sll.head, sll.tail = sll.tail, sll.head
}
