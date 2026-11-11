package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	store := make(map[int]struct{})

	for _, num := range nums {
		store[num] = struct{}{}
	}

	longest := 0
	for num := range store {
		if _, exist := store[num-1]; !exist {
			length := 1
			for _, exist := store[num+length]; exist; _, exist = store[num+length] {
				length++
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}
