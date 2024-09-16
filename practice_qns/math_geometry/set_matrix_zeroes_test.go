package mathgeometry_test

import (
	"fmt"
	mathgeometry "go-data-structure/practice_qns/math_geometry"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		input [][]int
		want  [][]int
	}{
		{
			input: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			input: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			mathgeometry.SetZeroes(test.input)
			assert.ElementsMatch(t, test.want, test.input)
		})
	}
}
