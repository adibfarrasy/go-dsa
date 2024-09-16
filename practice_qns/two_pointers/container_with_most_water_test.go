package twopointers_test

import (
	"fmt"
	twopointers "go-data-structure/practice_qns/two_pointers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:  49,
		},
		{
			input: []int{1, 1},
			want:  1,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := twopointers.MaxArea(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
