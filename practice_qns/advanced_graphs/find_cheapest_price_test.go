package advancedgraphs_test

import (
	"fmt"
	advancedgraphs "go-data-structure/practice_qns/advanced_graphs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCheapestPrice(t *testing.T) {
	tests := []struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
		want    int
	}{
		{
			n: 4,
			flights: [][]int{
				{0, 1, 100},
				{1, 2, 100},
				{2, 0, 100},
				{1, 3, 600},
				{2, 3, 200},
			},
			src:  0,
			dst:  3,
			k:    1,
			want: 700,
		},
		{
			n: 3,
			flights: [][]int{
				{0, 1, 100},
				{1, 2, 100},
				{0, 2, 500},
			},
			src:  0,
			dst:  2,
			k:    1,
			want: 200,
		},
		{
			n: 3,
			flights: [][]int{
				{0, 1, 100},
				{1, 2, 100},
				{0, 2, 500},
			},
			src:  0,
			dst:  2,
			k:    0,
			want: 500,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := advancedgraphs.FindCheapestPrice(test.n, test.flights, test.src, test.dst, test.k)
			assert.Equal(t, test.want, got)
		})
	}
}
