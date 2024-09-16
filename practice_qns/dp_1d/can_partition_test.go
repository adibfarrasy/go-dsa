package dp1d_test

import (
	"fmt"
	dp1d "go-data-structure/practice_qns/dp_1d"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartition(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 5, 11, 5},
			want: true,
		},
		{
			nums: []int{1, 2, 3, 5},
			want: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := dp1d.CanPartition(test.nums)
			assert.Equal(t, test.want, got)
		})
	}
}
