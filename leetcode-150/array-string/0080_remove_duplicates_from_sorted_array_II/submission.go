package removeduplicatesfromsortedarrayii

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	k := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
