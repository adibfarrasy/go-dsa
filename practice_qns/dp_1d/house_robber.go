package dp1d

func Rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[n-1] = nums[n-1]
	dp[n-2] = nums[n-2]

	for i := n - 3; i >= 0; i-- {
		dp[i] = nums[i] + max(dp[i+2], dp[i+3])
	}

	return max(dp[0], dp[1])
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
