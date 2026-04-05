func hasDuplicate(nums []int) bool {
	dupe := false
    dict := make(map[int]bool)

	for _, n := range nums {
		if dict[n] {
			dupe = true
			break
		}
		dict[n] = true
	}
	return dupe
}
