package stack

import "fmt"

type Stack[T any] struct {
	top    *node[T]
	length uint
}

type node[T any] struct {
	value T
	next  *node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	if s == nil {
		panic("Stack is not initialized")
	}

	s.length++
	newNode := &node[T]{value: value, next: s.top}
	if s.length == 1 {
		s.top = newNode
		return
	}
	s.top = newNode
}

func (s *Stack[T]) Pop() T {
	if s == nil {
		panic("Stack is not initialized")
	} else if s.length == 0 {
		panic("Pop from Empty Stack")
	}

	s.length--
	top := s.top
	if s.length == 0 {
		s.top = nil
		return top.value
	}

	s.top = s.top.next
	top.next = nil
	return top.value
}

func (s *Stack[T]) Print() {
	if s == nil {
		panic("Stack is not initialized")
	} else if s.length == 0 {
		return
	}

	fmt.Print(s.top.value)
	curr := s.top.next
	for curr != nil {
		fmt.Printf(" <- %v", curr.value)
		curr = curr.next
	}
	fmt.Println()
}

func (s *Stack[T]) Length() uint {
	if s == nil {
		panic("Stack is not initialized")
	}
	return s.length
}

func (s *Stack[T]) Peek() T {
	if s == nil {
		panic("Stack is not initialized")
	} else if s.length == 0 {
		panic("Peek from Empty Stack")
	}

	return s.top.value
}

func (s *Stack[T]) Clear() {
	if s == nil {
		panic("Stack is not initialized")
	}

	s.top = nil
	s.length = 0
}
