package backtracking

func Subsets(nums []int) [][]int {
	res := [][]int{}

	curr := []int{}
	var dfsHelper func(idx int)
	dfsHelper = func(idx int) {
		currCopy := append([]int{}, curr...)
		res = append(res, currCopy)
		if idx == len(nums) {
			return
		}

		for i := idx; i < len(nums); i++ {
			curr = append(curr, nums[i])
			dfsHelper(i + 1)
			curr = curr[:len(curr)-1]
		}
	}

	dfsHelper(0)
	return res
}
