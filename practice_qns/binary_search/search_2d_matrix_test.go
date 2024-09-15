package binarysearch_test

import (
	"fmt"
	binarysearch "go-data-structure/practice_qns/binary_search"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 3,
			want:   true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := binarysearch.SearchMatrix(test.matrix, test.target)
			assert.Equal(t, test.want, got)
		})
	}
}
