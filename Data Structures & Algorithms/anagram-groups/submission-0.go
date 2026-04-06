func groupAnagrams(strs []string) [][]string {
	var solution [][]string
	table := make(map[[26]int][]string)

	for _, str := range strs {
		m := countLetters(str)
		table[m] = append(table[m], str)
	}
	for _, v := range table {
		solution = append(solution, v)
	}

	return solution
}

func countLetters(s string) [26]int {
	smap := [26]int{}
	
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		smap[c] += 1
	}

	return smap
}
