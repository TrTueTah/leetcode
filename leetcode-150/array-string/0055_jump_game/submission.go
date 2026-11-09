package jumpgame

func canJump(nums []int) bool {
	maxJump := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxJump < i {
			return false
		}
		if nums[i]+i > maxJump {
			maxJump = nums[i] + i
		}
	}
	return maxJump >= len(nums)-1
}
