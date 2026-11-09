package jumpgameii

func jump(nums []int) int {
	count, currFarthest, nextFarthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		nextFarthest = max(nextFarthest, nums[i]+i)
		if currFarthest == i {
			count++
			currFarthest = nextFarthest
		}
	}
	return count
}
