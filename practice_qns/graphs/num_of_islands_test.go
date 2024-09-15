package graphs_test

import (
	"fmt"
	"go-data-structure/practice_qns/graphs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReoderList(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := graphs.NumIslands(test.grid)
			assert.Equal(t, test.want, got)
		})
	}
}
