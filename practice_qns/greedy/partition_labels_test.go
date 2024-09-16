package greedy_test

import (
	"fmt"
	"go-data-structure/practice_qns/greedy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		s    string
		want []int
	}{
		{
			s:    "ababcbacadefegdehijhklij",
			want: []int{9, 7, 8},
		},
		{
			s:    "eccbbbbdec",
			want: []int{10},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := greedy.PartitionLabels(test.s)
			assert.Equal(t, test.want, got)
		})
	}
}
