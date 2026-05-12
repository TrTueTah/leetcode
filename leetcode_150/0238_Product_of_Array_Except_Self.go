func productExceptSelf(nums []int) []int {
	n := len(nums)
    pre := make([]int, n)
	suf := make([]int, n)
	pre[0] = nums[0]
	suf[len(nums) - 1] = nums[len(nums) - 1]
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i - 1] * nums[i]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		suf[i] = suf[i + 1] * nums[i]
	}
	for i := 1; i < len(nums) - 1; i++ {
		nums[i] = pre[i - 1] * suf[i + 1]
	}
	nums[0] = suf[1]
	nums[len(nums) - 1] = pre[len(nums) - 2]
	return nums
}