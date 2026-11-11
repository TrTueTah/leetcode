package twosum

func twoSum(nums []int, target int) []int {
	store := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if val, exist := store[target-nums[i]]; exist {
			return []int{val, i}
		}
		store[nums[i]] = i
	}
	return []int{}
}
