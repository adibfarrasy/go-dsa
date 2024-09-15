package heap_test

import (
	"fmt"
	"go-data-structure/practice_qns/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReoderList(t *testing.T) {
	tests := []struct {
		points [][]int
		k      int
		want   [][]int
	}{
		{
			points: [][]int{
				{3, 3},
				{5, -1},
				{-2, 4},
			},
			k: 2,
			want: [][]int{
				{-2, 4},
				{3, 3},
			},
		},
		{
			points: [][]int{
				{1, 3},
				{-2, 2},
			},
			k: 1,
			want: [][]int{
				{-2, 2},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := heap.KClosest(test.points, test.k)
			assert.ElementsMatch(t, test.want, got)
		})
	}
}
