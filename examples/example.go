package main

import (
	"cmp"
	"fmt"

	dsg "github.com/vrv501/ds-go/dll"

	ssg "github.com/vrv501/ds-go/sll"

	"github.com/vrv501/ds-go/stack"

	"github.com/vrv501/ds-go/queue"
)

func main() {
	fmt.Println("Doubly linked list")
	dll := dsg.NewDoublyLinkedList(cmp.Compare[int])
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)
	dll.Prepend(0)
	dll.InsertAtIndex(1, 1)
	dll.Print()
	dll.Reverse()
	dll.Print()

	fmt.Println("Singly linked list")
	sll := ssg.NewSinglyLinkedList(cmp.Compare[int])
	sll.Append(1)
	sll.Append(2)
	sll.Append(3)
	sll.Prepend(0)
	sll.InsertAtIndex(1, 1)
	sll.Print()
	sll.Reverse()
	sll.Print()

	fmt.Println("Stack")
	st := stack.NewStack[int]()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Print()
	st.Pop()
	st.Print()
	fmt.Println(st.Peek())
	st.Clear()

	fmt.Println("Queue")
	q := queue.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Print()
	q.Dequeue()
	q.Print()
	fmt.Println(q.Peek())
	q.Clear()
}
