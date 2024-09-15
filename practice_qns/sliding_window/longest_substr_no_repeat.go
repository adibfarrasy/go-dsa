package slidingwindow

func LengthOfLongestSubstring(s string) int {
	charset := map[byte]bool{}
	res := 0

	l := 0
	for r := range s {
		for charset[s[r]] {
			delete(charset, s[l])
			l++
		}
		charset[s[r]] = true
		res = max(res, r-l+1)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
