package slidingwindow

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	charset := [26]int{}
	for _, c := range s1 {
		charset[c-'a'] += 1
	}

	l, r := 0, len(s1)
	for r < len(s2) {
		substr := s2[l:r]

		tmpCharset := [26]int{}
		for _, c := range substr {
			tmpCharset[c-'a'] += 1
		}

		count := 0
		for i := range charset {
			if charset[i] == tmpCharset[i] {
				count++
			}
		}

		if count == 26 {
			return true
		}

		l += 1
		r += 1
	}

	return false
}
