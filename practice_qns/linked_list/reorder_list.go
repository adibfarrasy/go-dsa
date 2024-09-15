package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Construct(vals []int) *ListNode {
	var head *ListNode

	for i := len(vals) - 1; i >= 0; i-- {
		newNode := &ListNode{
			Val:  vals[i],
			Next: head,
		}
		head = newNode
	}

	return head
}

func (l *ListNode) String() string {
	res := []int{}

	curr := l
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}

	return fmt.Sprintf("%+v", res)
}

func ReorderList(head *ListNode) {
	// find the middle of the list
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	reversed := reverse(slow.Next)
	slow.Next = nil

	curr := head

	for curr != nil && reversed != nil {
		next := curr.Next
		revNext := reversed.Next
		curr.Next = reversed
		reversed.Next = next
		curr = next
		reversed = revNext
	}
}

func reverse(node *ListNode) *ListNode {
	var prev, curr *ListNode = nil, node

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
