package ds_test

import (
	"fmt"
	"go-data-structure/ds"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	tests := []struct {
		name        string
		fn          func() (string, error)
		expectedRes string
		expectedErr error
	}{
		{
			name: "test DFS",
			fn: func() (string, error) {
				graph := ds.NewGraph[int](false)
				graph.AddEdge(1, 2)
				graph.AddEdge(2, 3)
				graph.AddEdge(1, 4)
				graph.AddEdge(3, 5)

				res := []int{}
				ds.DFS(&graph, 1, func(data int) {
					res = append(res, data)
				})

				return fmt.Sprintf("%+v", res), nil
			},
			expectedRes: "[1 2 3 5 4]",
		},
		{
			name: "test BFS",
			fn: func() (string, error) {
				graph := ds.NewGraph[int](false)
				graph.AddEdge(1, 2)
				graph.AddEdge(2, 3)
				graph.AddEdge(1, 4)
				graph.AddEdge(3, 5)

				res := []int{}
				ds.BFS(&graph, 1, func(data int) {
					res = append(res, data)
				})

				return fmt.Sprintf("%+v", res), nil
			},
			expectedRes: "[1 2 4 3 5]",
		},
	}

	for _, test := range tests {
		res, err := test.fn()
		assert.Equal(t, test.expectedErr, err)
		assert.Equal(t, test.expectedRes, res)
	}
}
