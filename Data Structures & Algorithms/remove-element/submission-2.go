func removeElement(nums []int, val int) int {

// inefficient
    // for i := 0; i < len(nums); {
	// 	if nums[i] == val {
	// 		nums = append(nums[:i], nums[i+1:]...)
	// 		continue
	// 	}
	// 	i++
	// }
	

	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
