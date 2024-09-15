package dp2d_test

import (
	"fmt"
	dp2d "go-data-structure/practice_qns/dp_2d"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m    int
		n    int
		want int
	}{
		{
			m:    3,
			n:    7,
			want: 28,
		},
		{
			m:    3,
			n:    2,
			want: 3,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := dp2d.UniquePaths(test.m, test.n)
			assert.Equal(t, test.want, got)
		})
	}
}
