package intervals

import (
	"sort"
)

func EraseOverlapIntervals(intervals [][]int) int {
	res := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prevEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		currInterval := intervals[i]
		if currInterval[0] < prevEnd {
			res += 1
			prevEnd = min(currInterval[1], prevEnd)
		} else {
			prevEnd = currInterval[1]
		}
	}

	return res
}
