package queue

import "fmt"

type Queue[T any] struct {
	first  *node[T]
	last   *node[T]
	length uint
}

type node[T any] struct {
	value T
	next  *node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	if q == nil {
		panic("Queue is not initialized")
	}

	q.length++
	newNode := &node[T]{value: value}
	if q.length == 1 {
		q.first = newNode
		q.last = newNode
		return
	}
	q.last.next = newNode
	q.last = newNode
}

func (q *Queue[T]) Dequeue() T {
	if q == nil {
		panic("Queue is not initialized")
	} else if q.length == 0 {
		panic("Dequeue from Empty Queue")
	}

	q.length--
	first := q.first
	if q.length == 0 {
		q.first = nil
		q.last = nil
		return first.value
	}

	q.first = q.first.next
	first.next = nil
	return first.value
}

func (q *Queue[T]) Print() {
	if q == nil {
		panic("Queue is not initialized")
	} else if q.length == 0 {
		return
	}

	fmt.Print(q.first.value)
	curr := q.first.next
	for curr != nil {
		fmt.Printf(" -> %v", curr.value)
		curr = curr.next
	}
	fmt.Println()
}

func (q *Queue[T]) Length() uint {
	if q == nil {
		panic("Queue is not initialized")
	}
	return q.length
}

func (q *Queue[T]) Peek() T {
	if q == nil {
		panic("Queue is not initialized")
	} else if q.length == 0 {
		panic("Peek from Empty Queue")
	}

	return q.first.value
}

func (q *Queue[T]) Clear() {
	if q == nil {
		panic("Queue is not initialized")
	}
	q.first = nil
	q.last = nil
	q.length = 0
}
