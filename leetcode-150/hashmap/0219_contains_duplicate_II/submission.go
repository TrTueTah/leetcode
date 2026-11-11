package containsduplicateii

func containsNearbyDuplicate(nums []int, k int) bool {
	store := make(map[int]int)
	for index, num := range nums {
		if val, exists := store[num]; exists && abs(val, index) <= k {
			return true
		}
		store[num] = index
	}

	return false
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
