package dp1d_test

import (
	"fmt"
	dp1d "go-data-structure/practice_qns/dp_1d"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHouseRobber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := dp1d.Rob(test.nums)
			assert.Equal(t, test.want, got)
		})
	}
}
