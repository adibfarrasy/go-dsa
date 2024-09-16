package intervals_test

import (
	"fmt"
	"go-data-structure/practice_qns/intervals"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      int
	}{
		{
			intervals: [][]int{
				{1, 2}, {2, 3}, {3, 4}, {1, 3},
			},
			want: 1,
		},
		{
			intervals: [][]int{
				{1, 2}, {1, 2}, {1, 2},
			},
			want: 2,
		},
		{
			intervals: [][]int{
				{1, 2}, {2, 3},
			},
			want: 0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := intervals.EraseOverlapIntervals(test.intervals)
			assert.Equal(t, test.want, got)
		})
	}
}
