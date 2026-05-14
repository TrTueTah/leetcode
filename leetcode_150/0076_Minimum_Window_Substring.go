func minWindow(s string, t string) string {
    if len(t) > len(s) {
		return ""
	}

	need := map[byte]int{}

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	required := len(need)
	formed := 0

	window := map[byte]int{}

	left := 0

	minLen := len(s) + 1
	start := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]

		window[ch]++

		if need[ch] > 0 && window[ch] == need[ch] {
			formed++
		}

		for formed == required {
			windowLen := right - left + 1

			if windowLen < minLen {
				minLen = windowLen
				start = left
			}

			leftChar := s[left]
			window[leftChar]--

			if need[leftChar] > 0 && window[leftChar] < need[leftChar] {
				formed--
			}

			left++
		}
	}

	if minLen == len(s) + 1 {
		return ""
	}
	return s[start : start + minLen]
}