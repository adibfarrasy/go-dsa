package ds

import "golang.org/x/exp/constraints"

type MinHeap[T constraints.Ordered] []T

func (h MinHeap[T]) Len() int { return len(h) }

func (h MinHeap[T]) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap[T]) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap[T]) Push(x any) { *h = append(*h, x.(T)) }

func (h *MinHeap[T]) Pop() any {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}
