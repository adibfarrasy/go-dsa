package ds_test

import (
	"fmt"
	"go-data-structure/ds"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {
	tests := []struct {
		name        string
		fn          func() (string, error)
		expectedRes string
		expectedErr error
	}{
		{
			name: "test in-order traversal",
			fn: func() (string, error) {
				tree := ds.BinaryTreeNode[int]{
					Val: 5,
				}
				tree.Insert(2)
				tree.Insert(10)

				res := []int{}
				ds.InOrder(&tree, func(data int) {
					res = append(res, data)
				})
				return fmt.Sprintf("%+v", res), nil
			},
			expectedRes: "[2 5 10]",
		},
		{
			name: "test pre-order traversal",
			fn: func() (string, error) {
				tree := ds.BinaryTreeNode[int]{
					Val: 5,
				}
				tree.Insert(2)
				tree.Insert(10)

				res := []int{}
				ds.PreOrder(&tree, func(data int) {
					res = append(res, data)
				})
				return fmt.Sprintf("%+v", res), nil
			},
			expectedRes: "[5 2 10]",
		},
		{
			name: "test post-order traversal",
			fn: func() (string, error) {
				tree := ds.BinaryTreeNode[int]{
					Val: 5,
				}
				tree.Insert(2)
				tree.Insert(10)

				res := []int{}
				ds.PostOrder(&tree, func(data int) {
					res = append(res, data)
				})
				return fmt.Sprintf("%+v", res), nil
			},
			expectedRes: "[2 10 5]",
		},
		{
			name: "test level-order traversal",
			fn: func() (string, error) {
				tree := ds.BinaryTreeNode[int]{
					Val: 5,
				}
				tree.Insert(2)
				tree.Insert(10)
				tree.Insert(-10)
				tree.Insert(4)
				tree.Insert(7)
				tree.Insert(20)

				res := []int{}
				ds.LevelOrder(&tree, func(data int) {
					res = append(res, data)
				})
				return fmt.Sprintf("%+v", res), nil
			},
			expectedRes: "[5 2 10 -10 4 7 20]",
		},
	}

	for _, test := range tests {
		res, err := test.fn()
		assert.Equal(t, test.expectedErr, err)
		assert.Equal(t, test.expectedRes, res)
	}
}
