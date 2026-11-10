package containerwithmostwater

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0
	for left < right {
		cur := min(height[left], height[right]) * (right - left)
		if cur > result {
			result = cur
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}
