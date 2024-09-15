package trees_test

import (
	"fmt"
	"go-data-structure/practice_qns/trees"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		input []any
		p     int
		q     int
		want  int
	}{
		{
			input: []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5},
			p:     2,
			q:     8,
			want:  6,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			root := trees.ConstructBST(test.input)
			p := root.GetByVal(test.p)
			q := root.GetByVal(test.q)

			got := trees.LowestCommonAncestor(root, p, q)
			if got != nil {
				assert.Equal(t, test.want, got.Val)
			} else {
				t.Fatalf("result not found")
			}
		})
	}
}
