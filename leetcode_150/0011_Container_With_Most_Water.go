func maxArea(height []int) int {
    left, right := 0, len(height) - 1
	result := 0
	for left < right {
		tmp := (right - left) * min(height[left], height[right])
		if tmp > result {
			result = tmp
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return result
}