package dll

import "fmt"

type DoublyLinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length uint
	// For most primitive types(integers, decimal, string), cmp.Compare can be used
	orderedFunc func(a, b T) int // should return 0 if equal, -ve value if a < b, and +ve if a > b
}

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

func NewDoublyLinkedList[T any](f func(a, b T) int) *DoublyLinkedList[T] {
	if f == nil {
		panic("Comparison function cannot be nil")
	}
	return &DoublyLinkedList[T]{orderedFunc: f}
}

func (dll *DoublyLinkedList[T]) Append(value T) {
	if dll == nil || dll.orderedFunc == nil {
		panic("Doubly Linked List is not initialized")
	}

	dll.length++
	newNode := &node[T]{value: value, prev: dll.tail}
	if dll.length == 1 {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	dll.tail.next = newNode
	dll.tail = newNode
}

func (dll *DoublyLinkedList[T]) Prepend(value T) {
	if dll == nil || dll.orderedFunc == nil {
		panic("Doubly Linked List is not initialized")
	}

	dll.length++
	newNode := &node[T]{value: value, next: dll.head}
	if dll.length == 1 {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	dll.head.prev = newNode
	dll.head = newNode
}

// Indexing starts at 0. If index > DLL length, function panics
func (dll *DoublyLinkedList[T]) InsertAtIndex(value T, index uint) {
	if dll == nil || dll.orderedFunc == nil {
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
	for index > 0 {
		curr = curr.next
		index--
	}

	dll.length++
	newNode := &node[T]{value: value, next: curr, prev: curr.prev}
	curr.prev.next = newNode
	curr.prev = newNode
}

func (dll *DoublyLinkedList[T]) Print() {
	if dll == nil || dll.orderedFunc == nil {
		panic("Doubly Linked List is not initialized")
	} else if dll.length == 0 {
		return
	}

	fmt.Print(dll.head.value)
	curr := dll.head.next
	for curr != nil {
		fmt.Printf(" <-> %v", curr.value)
		curr = curr.next
	}
	fmt.Println()
}

func (dll *DoublyLinkedList[T]) Length() uint {
	if dll == nil || dll.orderedFunc == nil {
		panic("Doubly Linked List is not initialized")
	}

	return dll.length
}

func (dll *DoublyLinkedList[T]) Pop() T {
	if dll == nil || dll.orderedFunc == nil {
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

func (dll *DoublyLinkedList[T]) PopFromFront() T {
	if dll == nil || dll.orderedFunc == nil {
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

// Returns value at index after removing it. If index > length or DLL length == 0, function panics
func (dll *DoublyLinkedList[T]) RemoveAtIndex(index uint) T {
	if dll == nil || dll.orderedFunc == nil {
		panic("Doubly Linked List is not initialized")
	} else if dll.length == 0 {
		panic("RemoveAtIndex From Empty Doubly Linked List")
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

// Returns true when value is found and removed
func (dll *DoublyLinkedList[T]) RemoveVal(val T) bool {
	if dll == nil || dll.orderedFunc == nil {
		panic("Doubly Linked List is not initialized")
	} else if dll.length == 0 {
		return false
	}

	curr := dll.head
	for curr != nil && dll.orderedFunc(val, curr.value) != 0 {
		curr = curr.next
	}

	switch {
	case curr == nil:
		return false
	case curr.prev == nil:
		_ = dll.PopFromFront()
		return true
	case curr.next == nil:
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

// Returns index if value is found
func (dll *DoublyLinkedList[T]) Contains(val T) (uint, bool) {
	if dll == nil || dll.orderedFunc == nil {
		panic("Doubly Linked List is not initialized")
	} else if dll.length == 0 {
		return 0, false
	}

	curr := dll.head
	var index uint
	for curr != nil && dll.orderedFunc(val, curr.value) != 0 {
		curr = curr.next
		index++
	}

	if curr == nil {
		return 0, false
	}
	return index, true
}

func (dll *DoublyLinkedList[T]) Reverse() {
	if dll == nil || dll.orderedFunc == nil {
		panic("Doubly Linked List is not initialized")
	} else if dll.length <= 1 {
		return
	}

	curr := dll.head
	for curr != nil {
		curr.prev, curr.next = curr.next, curr.prev
		curr = curr.prev
	}
	dll.head, dll.tail = dll.tail, dll.head
}
