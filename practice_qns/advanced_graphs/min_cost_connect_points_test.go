package advancedgraphs_test

import (
	"fmt"
	advancedgraphs "go-data-structure/practice_qns/advanced_graphs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostConnectPoints(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{
			points: [][]int{
				{0, 0},
				{2, 2},
				{3, 10},
				{5, 2},
				{7, 0},
			},
			want: 20,
		},
		{
			points: [][]int{
				{3, 12},
				{-2, 5},
				{-4, 1},
			},
			want: 18,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := advancedgraphs.MinCostConnectPoints(test.points)
			assert.Equal(t, test.want, got)
		})
	}
}
