package ds_test

import (
	"go-data-structure/ds"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name        string
		fn          func() (int, error)
		expectedRes int
		expectedErr error
	}{
		{
			name: "pop empty stack",
			fn: func() (int, error) {
				stack := ds.NewStack[int]()
				return stack.Pop()
			},
			expectedRes: 0,
			expectedErr: ds.ErrPopEmpty,
		},
		{
			name: "pop last inserted",
			fn: func() (int, error) {
				stack := ds.NewStack[int]()
				stack.Push(1)
				stack.Push(2)
				stack.Push(3)
				return stack.Pop()
			},
			expectedRes: 3,
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		res, err := test.fn()
		assert.Equal(t, test.expectedErr, err)
		assert.Equal(t, test.expectedRes, res)
	}
}
