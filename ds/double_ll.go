package ds

import "fmt"

type DNode[T comparable] struct {
	Val  T
	Prev *DNode[T]
	Next *DNode[T]
}

type DoublyLinkedList[T comparable] struct {
	head *DNode[T]
	tail *DNode[T]
}

func NewDoublyLinkedList[T comparable]() DoublyLinkedList[T] {
	return DoublyLinkedList[T]{}
}

func (ll *DoublyLinkedList[T]) InsertAtStart(v T) {
	newNode := &DNode[T]{Val: v}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.Next = ll.head
		ll.head.Prev = newNode
		ll.head = newNode
	}

}

func (ll *DoublyLinkedList[T]) InsertAtEnd(v T) {
	newNode := &DNode[T]{Val: v}

	if ll.tail == nil {
		ll.tail = newNode
		ll.head = newNode
	} else {
		newNode.Prev = ll.tail
		ll.tail.Next = newNode
		ll.tail = newNode
	}
}

func (ll *DoublyLinkedList[T]) Delete(v T) {
	curr := ll.head
	for curr != nil && curr.Val != v {
		curr = curr.Next
	}

	if curr == nil {
		return
	}

	if curr.Prev != nil {
		curr.Prev.Next = curr.Next
	} else {
		ll.head = curr.Next
	}

	if curr.Next != nil {
		curr.Next.Prev = curr.Prev
	} else {
		ll.tail = curr.Prev
	}
}

func (ll *DoublyLinkedList[T]) String() string {
	res := ""
	curr := ll.head
	for curr != nil {
		res = res + fmt.Sprintf("%v -> ", curr.Val)
		curr = curr.Next
	}
	res = res + fmt.Sprint("nil")
	return res
}
