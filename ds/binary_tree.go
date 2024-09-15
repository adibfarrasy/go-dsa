package ds

import (
	"golang.org/x/exp/constraints"
)

type BinaryTreeNode[T constraints.Ordered] struct {
	Val   T
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}

func (t *BinaryTreeNode[T]) Insert(v T) {
	newNode := &BinaryTreeNode[T]{Val: v}

	if v <= t.Val {
		if t.Left == nil {
			t.Left = newNode
		} else {
			t.Left.Insert(v)
		}
	} else {
		if t.Right == nil {
			t.Right = newNode
		} else {
			t.Right.Insert(v)
		}
	}
}

func InOrder[T constraints.Ordered](node *BinaryTreeNode[T], visit func(data T)) {
	if node == nil {
		return
	}

	InOrder(node.Left, visit)
	visit(node.Val)
	InOrder(node.Right, visit)
}

func PreOrder[T constraints.Ordered](node *BinaryTreeNode[T], visit func(data T)) {
	if node == nil {
		return
	}

	visit(node.Val)
	PreOrder(node.Left, visit)
	PreOrder(node.Right, visit)
}

func PostOrder[T constraints.Ordered](node *BinaryTreeNode[T], visit func(data T)) {
	if node == nil {
		return
	}

	PreOrder(node.Left, visit)
	PreOrder(node.Right, visit)
	visit(node.Val)
}

func LevelOrder[T constraints.Ordered](node *BinaryTreeNode[T], visit func(data T)) {
	if node == nil {
		return
	}

	queue := NewQueue[*BinaryTreeNode[T]]()
	queue.Push(node)

	for queue.Len() > 0 {
		curr, _ := queue.Pop()
		visit(curr.Val)

		if curr.Left != nil {
			queue.Push(curr.Left)
		}

		if curr.Right != nil {
			queue.Push(curr.Right)
		}
	}
}
