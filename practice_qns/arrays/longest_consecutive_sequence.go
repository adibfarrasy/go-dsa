package arrays

func LongestConsecutive(nums []int) int {
	set := map[int]bool{}

	for _, num := range nums {
		set[num] = true
	}

	res := 0

	for _, num := range nums {
		if set[num-1] {
			continue
		}

		seq := 1
		temp := num + 1

		for set[temp] {
			seq += 1
			temp += 1
		}

		if seq > res {
			res = seq
		}
	}
	return res
}
