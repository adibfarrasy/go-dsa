package arrays_test

import (
	"fmt"
	"go-data-structure/practice_qns/arrays"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{100, 4, 200, 1, 3, 2},
			want:  4,
		},
		{
			input: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want:  9,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := arrays.LongestConsecutive(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
