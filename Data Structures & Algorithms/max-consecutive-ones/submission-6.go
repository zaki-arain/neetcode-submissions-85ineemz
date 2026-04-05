func findMaxConsecutiveOnes(nums []int) int {
    final := 0
    count := 0

    for i := 0; i < len(nums); i++ {
        if (nums[i] == 1) {
            count++
        } else {
            final = max(final, count)
            count = 0
        }
    }
    return max(final, count)
}