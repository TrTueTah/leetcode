package sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for k := 0; k < len(nums)-2; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		left, right := k+1, len(nums)-1
		for left < right {
			sum := nums[k] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[k], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
