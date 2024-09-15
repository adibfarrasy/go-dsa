package mathgeometry_test

import (
	"fmt"
	mathgeometry "go-data-structure/practice_qns/math_geometry"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateImage(t *testing.T) {
	tests := []struct {
		input [][]int
		want  [][]int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			mathgeometry.Rotate(test.input)
			assert.ElementsMatch(t, test.want, test.input)
		})
	}
}
