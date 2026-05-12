func jump(nums []int) int {
	n := len(nums)
	jumps, current_end, farthest := 0, 0, 0
	for i := 0; i < n - 1; i++ {
		farthest = max(farthest, nums[i] + i)
		if i == current_end {
			++jumps
			current_end = farthest
		}
	}
	return jumps
}