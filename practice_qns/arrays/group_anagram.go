package arrays

func GroupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	for _, str := range strs {
		count := [26]int{}
		for _, c := range str {
			count[c-'a'] += 1
		}
		anagramMap[count] = append(anagramMap[count], str)
	}

	result := [][]string{}
	for _, anagrams := range anagramMap {
		result = append(result, anagrams)
	}

	return result
}
