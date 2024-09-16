package stack_test

import (
	"fmt"
	"go-data-structure/practice_qns/stack"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarFleet(t *testing.T) {
	tests := []struct {
		target   int
		position []int
		speed    []int
		want     int
	}{
		{
			target:   12,
			position: []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			want:     3,
		},
		{
			target:   10,
			position: []int{3},
			speed:    []int{3},
			want:     1,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := stack.CarFleet(test.target, test.position, test.speed)
			assert.Equal(t, test.want, got)
		})
	}
}
