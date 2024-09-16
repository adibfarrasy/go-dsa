package mathgeometry

func SetZeroes(matrix [][]int) {
	ROWS, COLS := len(matrix), len(matrix[0])
	zeroRows := make([]bool, ROWS)
	zeroCols := make([]bool, COLS)

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if matrix[i][j] == 0 {
				zeroRows[i] = true
				zeroCols[j] = true
			}
		}
	}

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if zeroRows[i] || zeroCols[j] {
				matrix[i][j] = 0
			}
		}
	}
}
