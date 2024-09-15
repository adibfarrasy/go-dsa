package arrays_test

import (
	"fmt"
	"go-data-structure/practice_qns/arrays"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input []string
		want  [][]string
	}{
		{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"},
				{"eat", "tea", "ate"},
				{"tan", "nat"},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := arrays.GroupAnagrams(test.input)
			assert.ElementsMatch(t, test.want, got)
		})
	}
}
