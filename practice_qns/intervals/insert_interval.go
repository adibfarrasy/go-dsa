package intervals

func Insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}

	for i, currIntv := range intervals {
		if newInterval[1] < currIntv[0] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			return res
		} else if newInterval[0] > currIntv[1] {
			res = append(res, currIntv)
		} else {
			newInterval[0] = min(newInterval[0], currIntv[0])
			newInterval[1] = max(newInterval[1], currIntv[1])
		}
	}

	res = append(res, newInterval)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
