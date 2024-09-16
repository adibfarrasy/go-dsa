package slidingwindow_test

import (
	"fmt"
	slidingwindow "go-data-structure/practice_qns/sliding_window"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringPermutation(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{
			s1:   "ab",
			s2:   "eidbaooo",
			want: true,
		},
		{
			s1:   "ab",
			s2:   "eidboaooo",
			want: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := slidingwindow.CheckInclusion(test.s1, test.s2)
			assert.Equal(t, test.want, got)
		})
	}
}
