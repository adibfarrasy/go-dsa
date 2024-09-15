package slidingwindow_test

import (
	"fmt"
	slidingwindow "go-data-structure/practice_qns/sliding_window"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "abcabcbb",
			want:  3,
		},
		{
			input: "bbbbb",
			want:  1,
		},
		{
			input: "pwwkew",
			want:  3,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := slidingwindow.LengthOfLongestSubstring(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
