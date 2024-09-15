package stack_test

import (
	"go-data-structure/practice_qns/stack"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	tests := []func(t *testing.T){
		func(t *testing.T) {
			minStack := stack.Constructor()
			minStack.Push(-2)
			minStack.Push(0)
			minStack.Push(-3)

			got := minStack.GetMin()
			assert.Equal(t, -3, got)

			minStack.Pop()

			got = minStack.Top()
			assert.Equal(t, 0, got)

			got = minStack.GetMin()
			assert.Equal(t, -2, got)
		},
	}

	for _, test := range tests {
		test(t)
	}
}
