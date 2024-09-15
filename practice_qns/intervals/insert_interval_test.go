package intervals_test

import (
	"fmt"
	"go-data-structure/practice_qns/intervals"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertInterval(t *testing.T) {
	tests := []struct {
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			intervals: [][]int{
				{1, 3},
				{6, 9},
			},
			newInterval: []int{2, 5},
			want: [][]int{
				{1, 5},
				{6, 9},
			},
		},
		{
			intervals: [][]int{
				{1, 2},
				{3, 5},
				{6, 7},
				{8, 10},
				{12, 16},
			},
			newInterval: []int{4, 8},
			want: [][]int{
				{1, 2},
				{3, 10},
				{12, 16},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := intervals.Insert(test.intervals, test.newInterval)
			assert.ElementsMatch(t, test.want, got)
		})
	}
}
