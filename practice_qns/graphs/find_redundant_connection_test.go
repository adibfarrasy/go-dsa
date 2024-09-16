package graphs_test

import (
	"fmt"
	"go-data-structure/practice_qns/graphs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		edges [][]int
		want  []int
	}{
		{
			edges: [][]int{{1, 2}, {1, 3}, {2, 3}},
			want:  []int{2, 3},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := graphs.FindRedundantConnection(test.edges)
			assert.Equal(t, test.want, got)
		})
	}
}
