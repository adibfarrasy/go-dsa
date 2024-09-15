package mathgeometry

func Rotate(matrix [][]int) {
	left, right := 0, len(matrix)-1
	for left < right {
		for i := 0; i < right-left; i++ {
			top, bottom := left, right
			topleft := matrix[top][left+i]

			matrix[top][left+i] = matrix[bottom-i][left]
			matrix[bottom-i][left] = matrix[bottom][right-i]
			matrix[bottom][right-i] = matrix[top+i][right]
			matrix[top+i][right] = topleft
		}

		right -= 1
		left += 1
	}
}
