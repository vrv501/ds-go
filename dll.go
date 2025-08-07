package dsa

import "fmt"

type doublyLinkedList[T any] struct {
	head        *node[T]
	tail        *node[T]
	length      uint
	orderedFunc func(a, b T) int // should return 0 if equal, -ve value if a < b, and +ve if a > b
	// For most primitive types, cmp.Compare can be used
}

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

func NewDoublyLinkedList[T any](f func(a, b T) int) *doublyLinkedList[T] {
	return &doublyLinkedList[T]{orderedFunc: f}
}

func (dll *doublyLinkedList[T]) Append(value T) {
	if dll == nil {
		panic("Doubly Linked List is not initialized")
	}

	dll.length++
	newNode := &node[T]{value: value, prev: dll.tail}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	dll.tail.next = newNode
	dll.tail = newNode
}

func (dll *doublyLinkedList[T]) Prepend(value T) {
	if dll == nil {
		panic("Doubly Linked List is not initialized")
	}

	dll.length++
	newNode := &node[T]{value: value, next: dll.head}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	dll.head.prev = newNode
	dll.head = newNode
}

func (dll *doublyLinkedList[T]) InsertAtIndex(value T, index uint) {
	if dll == nil {
		panic("Doubly Linked List is not initialized")
	} else if index == 0 {
		dll.Prepend(value)
		return
	} else if index == dll.length {
		dll.Append(value)
		return
	} else if index > dll.length {
		panic("Index out of bounds")
	}

	curr := dll.head
	dll.length++
	for index > 0 {
		curr = curr.next
		index--
	}

	newNode := &node[T]{value: value, next: curr, prev: curr.prev}
	curr.prev.next = newNode
	curr.prev = newNode
}

func (dll *doublyLinkedList[T]) Print() {
	if dll == nil {
		panic("Doubly Linked List is not initialized")
	}

	curr := dll.head
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func (dll *doublyLinkedList[T]) Length() uint {
	if dll == nil {
		panic("Doubly Linked List is not initialized")
	}

	return dll.length
}

func (dll *doublyLinkedList[T]) Pop() T {
	if dll == nil {
		panic("Doubly Linked List is not initialized")
	} else if dll.length == 0 {
		panic("Pop from Empty Doubly Linked List")
	}

	dll.length--
	val := dll.tail.value
	if dll.length == 0 {
		dll.head = nil
		dll.tail = nil
		return val
	}

	dll.tail = dll.tail.prev
	dll.tail.next.prev = nil
	dll.tail.next = nil
	return val
}

func (dll *doublyLinkedList[T]) PopFromFront() T {
	if dll == nil {
		panic("Doubly Linked List is not initialized")
	} else if dll.length == 0 {
		panic("PopFromFront from Empty Doubly Linked List")
	}

	dll.length--
	val := dll.head.value
	if dll.length == 0 {
		dll.head = nil
		dll.tail = nil
		return val
	}

	dll.head = dll.head.next
	dll.head.prev.next = nil
	dll.head.prev = nil
	return val
}

func (dll *doublyLinkedList[T]) RemoveAtIndex(index uint) T {
	if dll == nil {
		panic("Doubly Linked List is not initialized")
	} else if dll.length == 0 {
		panic("Remove From Empty Doubly Linked List")
	} else if index == 0 {
		return dll.PopFromFront()
	} else if index == dll.length-1 {
		return dll.Pop()
	} else if index >= dll.length {
		panic("Index out of bounds")
	}

	curr := dll.head
	dll.length--
	for index > 0 {
		curr = curr.next
		index--
	}

	curr.prev.next = curr.next
	curr.next.prev = curr.prev
	curr.next = nil
	curr.prev = nil
	return curr.value
}

func (dll *doublyLinkedList[T]) RemoveVal(val T) bool {
	if dll == nil {
		panic("Doubly Linked List is not initialized")
	} else if dll.length == 0 {
		return false
	}

	curr := dll.head
	for curr != nil && dll.orderedFunc(val, curr.value) != 0 {
		curr = curr.next
	}
	if curr == nil {
		return false
	}

	if curr == dll.head {
		_ = dll.PopFromFront()
		return true
	} else if curr == dll.tail {
		_ = dll.Pop()
		return true
	}
	dll.length--
	curr.prev.next = curr.next
	curr.next.prev = curr.prev
	curr.next = nil
	curr.prev = nil
	return true
}
