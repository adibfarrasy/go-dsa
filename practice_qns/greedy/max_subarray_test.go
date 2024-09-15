package greedy_test

import (
	"fmt"
	"go-data-structure/practice_qns/greedy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHouseRobber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			nums: []int{1},
			want: 1,
		},
		{
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := greedy.MaxSubArray(test.nums)
			assert.Equal(t, test.want, got)
		})
	}
}
