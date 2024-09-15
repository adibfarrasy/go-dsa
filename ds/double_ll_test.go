package ds_test

import (
	"go-data-structure/ds"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList(t *testing.T) {
	tests := []struct {
		name        string
		fn          func() (string, error)
		expectedRes string
		expectedErr error
	}{
		{
			name: "test insert at start and end",
			fn: func() (string, error) {
				doublyLL := ds.NewDoublyLinkedList[int]()
				doublyLL.InsertAtEnd(10)
				doublyLL.InsertAtEnd(20)
				doublyLL.InsertAtEnd(30)
				doublyLL.InsertAtStart(5)
				return doublyLL.String(), nil
			},
			expectedRes: "5 -> 10 -> 20 -> 30 -> nil",
		},
		{
			name: "test delete",
			fn: func() (string, error) {
				doublyLL := ds.NewDoublyLinkedList[int]()
				doublyLL.InsertAtEnd(10)
				doublyLL.InsertAtEnd(20)
				doublyLL.InsertAtEnd(30)
				doublyLL.Delete(10)
				return doublyLL.String(), nil
			},
			expectedRes: "20 -> 30 -> nil",
		},
	}

	for _, test := range tests {
		res, err := test.fn()
		assert.Equal(t, test.expectedErr, err)
		assert.Equal(t, test.expectedRes, res)
	}
}
