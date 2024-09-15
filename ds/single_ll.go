package ds

import "fmt"

type SNode[T comparable] struct {
	Val  T
	Next *SNode[T]
}

type SinglyLinkedList[T comparable] struct {
	Head *SNode[T]
}

func NewSinglyLinkedList[T comparable]() SinglyLinkedList[T] {
	return SinglyLinkedList[T]{}
}

func (ll *SinglyLinkedList[T]) InsertAtEnd(v T) {
	newNode := &SNode[T]{Val: v}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	curr := ll.Head
	if curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func (ll *SinglyLinkedList[T]) InsertAtStart(v T) {
	newNode := &SNode[T]{Val: v}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	curr := ll.Head
	ll.Head = newNode
	newNode.Next = curr
}

func (ll *SinglyLinkedList[T]) Delete(v T) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Val == v {
		ll.Head = ll.Head.Next
		return
	}

	curr := ll.Head
	for curr.Next != nil && curr.Next.Val != v {
		curr = curr.Next
	}

	if curr.Next != nil {
		curr.Next = curr.Next.Next
	}
}

func (ll *SinglyLinkedList[T]) String() string {
	res := ""
	curr := ll.Head
	for curr != nil {
		res = res + fmt.Sprintf("%v -> ", curr.Val)
		curr = curr.Next
	}
	res = res + fmt.Sprint("nil")
	return res
}
