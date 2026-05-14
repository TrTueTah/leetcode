func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	prefix := make([]int, n + 1)

	for i := 0; i < n; i++ {
		prefix[i + 1] = prefix[i] + nums[i]
	}

	minLen := n + 1

	for i := 0; i <= n; i++ {
		needed := prefix[i] + target
		left, right := i + 1, n
		firstTrue := -1

		for left <= right {
			mid := left + (right - left) / 2
			if prefix[mid] >= needed {
				firstTrue = mid
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if firstTrue != -1 {
			length := firstTrue - i
			if length < minLen {
				minLen = length
			}
		}
	}

	if minLen == n + 1 {
		return 0
	}
	return minLen
}