package bitmanipulation_test

import (
	"fmt"
	bitmanipulation "go-data-structure/practice_qns/bit_manipulation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			input: 123,
			want:  321,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := bitmanipulation.Reverse(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
