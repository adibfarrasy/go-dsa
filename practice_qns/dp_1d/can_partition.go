package dp1d

func CanPartition(nums []int) bool {
	sum := sum(nums)

	if sum%2 != 0 {
		return false
	}

	dp := map[int]bool{}
	dp[0] = true
	target := sum / 2

	for i := len(nums) - 1; i >= 0; i-- {
		nextDP := map[int]bool{}
		for n := range dp {
			if n+nums[i] == target {
				return true
			}
			nextDP[n+nums[i]] = true
			nextDP[n] = true
		}

		dp = nextDP
	}

	return false
}

func sum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
