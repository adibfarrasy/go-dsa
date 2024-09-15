package twopointers

func TwoSumII(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			// meant to return index + 1
			return []int{i + 1, j + 1}
		}
		if numbers[i]+numbers[j] < target {
			i += 1
		} else {
			j -= 1
		}
	}

	return nil
}
