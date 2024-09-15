package linkedlist_test

import (
	"fmt"
	linkedlist "go-data-structure/practice_qns/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReoderList(t *testing.T) {
	tests := []struct {
		input []int
		want  string
	}{
		{
			input: []int{1, 2, 3, 4},
			want:  "[1 4 2 3]",
		},
		{
			input: []int{1, 2, 3, 4, 5},
			want:  "[1 5 2 4 3]",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			head := linkedlist.Construct(test.input)
			linkedlist.ReorderList(head)
			assert.Equal(t, test.want, head.String())
		})
	}
}
