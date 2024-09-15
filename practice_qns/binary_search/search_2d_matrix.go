package binarysearch

func SearchMatrix(matrix [][]int, target int) bool {
	ROWS := len(matrix)
	COLS := len(matrix[0])

	top, bot := 0, ROWS-1

	for top <= bot {
		row := (top + bot) / 2
		if target > matrix[row][COLS-1] {
			top = row + 1
		} else if target < matrix[row][0] {
			bot = row - 1
		} else {
			break
		}
	}

	row := (top + bot) / 2
	left, right := 0, COLS-1
	for left <= right {
		mid := (left + right) / 2
		if target == matrix[row][mid] {
			return true
		} else if target > matrix[row][mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
