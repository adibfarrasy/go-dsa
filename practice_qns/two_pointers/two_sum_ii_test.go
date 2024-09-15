package twopointers_test

import (
	"fmt"
	twopointers "go-data-structure/practice_qns/two_pointers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := twopointers.TwoSumII(test.numbers, test.target)
			assert.Equal(t, test.want, got)
		})
	}
}
