package trappingrainwater

func trap(height []int) int {
	leftHeights, rightHeights := make([]int, len(height)), make([]int, len(height))
	leftHeights[0] = height[0]
	rightHeights[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		leftHeights[i] = max(leftHeights[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightHeights[i] = max(rightHeights[i+1], height[i])
	}
	result := 0
	for i := 1; i < len(height)-1; i++ {
		if min(leftHeights[i], rightHeights[i]) > height[i] {
			result += min(leftHeights[i], rightHeights[i]) - height[i]
		}
	}
	return result
}
