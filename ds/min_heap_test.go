package ds_test

import (
	"container/heap"
	"fmt"
	"go-data-structure/ds"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	tests := []struct {
		name        string
		fn          func() (string, error)
		expectedRes string
		expectedErr error
	}{
		{
			name: "test heap",
			fn: func() (string, error) {
				minHeap := &ds.MinHeap[int]{}
				heap.Init(minHeap)

				heap.Push(minHeap, 10)
				heap.Push(minHeap, 1)
				heap.Push(minHeap, 4)
				heap.Push(minHeap, 5)

				res := []int{}
				for minHeap.Len() > 0 {
					res = append(res, heap.Pop(minHeap).(int))
				}
				return fmt.Sprintf("%+v", res), nil
			},
			expectedRes: "[1 4 5 10]",
		},
	}

	for _, test := range tests {
		res, err := test.fn()
		assert.Equal(t, test.expectedErr, err)
		assert.Equal(t, test.expectedRes, res)
	}
}
