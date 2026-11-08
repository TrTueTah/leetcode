package rotatearray

func rotate(nums []int, k int) {
	k = k % len(nums)
	swap(nums, 0, len(nums)-1)
	swap(nums, 0, k-1)
	swap(nums, k, len(nums)-1)
}

func swap(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
