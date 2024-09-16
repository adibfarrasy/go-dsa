package trees_test

import (
	"go-data-structure/practice_qns/trees"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	tests := []func(t *testing.T){
		func(t *testing.T) {
			preorder := []int{3, 9, 20, 15, 7}
			inorder := []int{9, 3, 15, 20, 7}
			root := trees.BuildTree(preorder, inorder)

			reconstructPreorder := root.Preorder()
			reconstructInorder := root.Inorder()

			assert.Equal(t, preorder, reconstructPreorder)
			assert.Equal(t, inorder, reconstructInorder)
		},
	}

	for _, test := range tests {
		test(t)
	}
}
