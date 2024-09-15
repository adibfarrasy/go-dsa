package heap

import (
	"container/heap"
	"math"
)

type minHeap [][]int

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	return calculateEuclideanDist(h[i][0], h[i][1]) < calculateEuclideanDist(h[j][0], h[j][1])
}
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func (h *minHeap) Push(v any) {
	*h = append(*h, v.([]int))
}

func KClosest(points [][]int, k int) [][]int {
	minHeap := &minHeap{}

	heap.Init(minHeap)
	for _, p := range points {
		heap.Push(minHeap, p)
	}

	res := [][]int{}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(minHeap).([]int))
	}

	return res
}

func calculateEuclideanDist(x, y int) float64 {
	return math.Sqrt(float64(x*x + y*y))
}
