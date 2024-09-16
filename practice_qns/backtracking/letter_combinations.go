package backtracking

func LetterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}

	letterMap := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"v", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var dfsHelper func(idx int, curr string)
	dfsHelper = func(idx int, curr string) {
		if len(curr) == len(digits) {
			res = append(res, curr)
			return
		}

		dig := digits[idx]
		str := letterMap[dig]

		for i := 0; i < len(str); i++ {
			curr += string(str[i])
			dfsHelper(idx+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	dfsHelper(0, "")
	return res
}
