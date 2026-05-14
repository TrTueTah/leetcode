func lengthOfLongestSubstring(s string) int {
    seen := map[byte]struct{}{}

	left := 0
	maxLength := 0

	for right := 0; right < len(s); right++ {
		for _, exists := seen[s[right]]; exists; _, exists = seen[s[right]] {
			delete(seen, s[left])
			left++
		}

		seen[s[right]] = struct{}{}

		windowLength := right - left + 1
		if windowLength > maxLength {
			maxLength = windowLength
		}
	}

	return maxLength
}