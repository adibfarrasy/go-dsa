package backtracking_test

import (
	"fmt"
	"go-data-structure/practice_qns/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: "23",
			want: []string{
				"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
			},
		},
		{
			input: "",
			want:  []string{},
		},
		{
			input: "2",
			want: []string{
				"a", "b", "c",
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := backtracking.LetterCombinations(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
