func twoSum(nums []int, target int) []int {
	var solution []int
	table := make(map[int]int) // number: index

	for i, n := range nums {
		diff := target - n

		diffIdx, isPresent := table[diff]
		if isPresent {
			solution = []int{diffIdx, i}
			break
		}

		if _, ok := table[n]; !ok {
			table[n] = i
		}
	}
	return solution
}
