func majorityElement(nums []int) int {
	store := map[int]int{}

	for i := 0; i < len(nums); i++ {
		store[nums[i]]++
	}

	maxCount := 0
	result := 0

	for key, value := range store {
		if value > maxCount {
			maxCount = value
			result = key
		}
	}

	return result
}