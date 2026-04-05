func isAnagram(s string, t string) bool {
	if len(s) != len(t) {return false}

	smap := make(map[string]int)
	tmap := make(map[string]int)
	for i, r := range s {
		smap[string(r)] += 1
		tmap[string(t[i])] += 1
	}
	return fmt.Sprint(smap) == fmt.Sprint(tmap)
}
