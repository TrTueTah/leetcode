package majorityelement

func majorityElement(nums []int) int {
	store := make(map[int]int)

	for _, num := range nums {
		store[num]++
	}
	max, res := 0, 0
	for key, val := range store {
		if val > max {
			res = key
			max = val
		}
	}
	return res
}
