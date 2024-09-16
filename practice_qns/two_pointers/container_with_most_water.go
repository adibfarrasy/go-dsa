package twopointers

func MaxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0

	for l < r {
		area := (r - l) * min(height[l], height[r])
		res = max(area, res)

		if height[l] < height[r] {
			l += 1
		} else {
			r -= 1
		}
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
