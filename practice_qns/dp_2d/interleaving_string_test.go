package dp2d_test

import (
	"fmt"
	dp2d "go-data-structure/practice_qns/dp_2d"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbcbcac",
			want: true,
		},
		{
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbbaccc",
			want: false,
		},
		{
			s1:   "",
			s2:   "",
			s3:   "",
			want: true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := dp2d.IsInterleave(test.s1, test.s2, test.s3)
			assert.Equal(t, test.want, got)
		})
	}
}
