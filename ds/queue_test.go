package ds_test

import (
	"go-data-structure/ds"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name        string
		fn          func() (int, error)
		expectedRes int
		expectedErr error
	}{
		{
			name: "pop empty queue",
			fn: func() (int, error) {
				queue := ds.NewQueue[int]()
				return queue.Pop()
			},
			expectedRes: 0,
			expectedErr: ds.ErrPopEmpty,
		},
		{
			name: "pop first inserted",
			fn: func() (int, error) {
				queue := ds.NewQueue[int]()
				queue.Push(1)
				queue.Push(2)
				queue.Push(3)
				return queue.Pop()
			},
			expectedRes: 1,
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		res, err := test.fn()
		assert.Equal(t, test.expectedErr, err)
		assert.Equal(t, test.expectedRes, res)
	}
}
