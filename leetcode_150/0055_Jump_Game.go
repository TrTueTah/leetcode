func canJump(nums []int) bool {
    reachable := 0
	for i := 0; i < len(nums); i++ {
		if reachable < i {
			return false;
		}
		reachable = max(reachable, nums[i] + i)
	}
	return true;
}