package bitmanipulation_test

import (
	"fmt"
	bitmanipulation "go-data-structure/practice_qns/bit_manipulation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSum(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{
			a:    1,
			b:    2,
			want: 3,
		},
		{
			a:    2,
			b:    3,
			want: 5,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := bitmanipulation.GetSum(test.a, test.b)
			assert.Equal(t, test.want, got)
		})
	}
}
