package ds_test

import (
	"go-data-structure/ds"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList(t *testing.T) {
	tests := []struct {
		name        string
		fn          func() (string, error)
		expectedRes string
		expectedErr error
	}{
		{
			name: "insert at start and end",
			fn: func() (string, error) {
				singeLL := ds.NewSinglyLinkedList[int]()
				singeLL.InsertAtStart(1)
				singeLL.InsertAtEnd(2)
				return singeLL.String(), nil
			},
			expectedRes: "1 -> 2 -> nil",
		},
		{
			name: "delete",
			fn: func() (string, error) {
				singeLL := ds.NewSinglyLinkedList[int]()
				singeLL.InsertAtStart(3)
				singeLL.InsertAtStart(2)
				singeLL.InsertAtStart(1)
				singeLL.Delete(2)
				return singeLL.String(), nil
			},
			expectedRes: "1 -> 3 -> nil",
		},
	}

	for _, test := range tests {
		res, err := test.fn()
		assert.Equal(t, test.expectedErr, err)
		assert.Equal(t, test.expectedRes, res)
	}
}
