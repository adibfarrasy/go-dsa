package backtracking_test

import (
	"fmt"
	"go-data-structure/practice_qns/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		input []int
		want  [][]int
	}{
		{
			input: []int{1, 2, 3},
			want: [][]int{
				{},
				{1},
				{2},
				{3},
				{1, 2},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
		{
			input: []int{0},
			want: [][]int{
				{},
				{0},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := backtracking.Subsets(test.input)
			assert.ElementsMatch(t, test.want, got)
		})
	}
}
