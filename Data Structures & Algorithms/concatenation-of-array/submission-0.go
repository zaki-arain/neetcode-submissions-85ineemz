func getConcatenation(nums []int) []int {
    ans := make([]int, len(nums) * 2)
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
		ans[i + len(nums)] = nums[i]
	}
	return ans
}


func test(nums []int) []int {
	var ans []int
	ans = append(ans, nums...)
	ans = append(ans, nums...)
	return ans
}