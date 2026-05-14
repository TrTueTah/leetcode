func findSubstring(s string, words []string) []int {
    if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	wordCount := map[string]int{}

	for _, word := range words {
		wordCount[word]++
	}

	result := []int{}

	stringLength := len(s)
	wordArrayLength := len(words)
	wordLength := len(words[0])

	for start := 0; start < wordLength; start++ {
		left, right := start, start

		currentWindowCount := map[string]int{}

		for right + wordLength <= stringLength {
			currentWord := s[right : right + wordLength]
			right += wordLength

			if _, exists := wordCount[currentWord]; !exists {
				currentWindowCount = map[string]int{}
				left = right
				continue
			}

			currentWindowCount[currentWord]++

			for currentWindowCount[currentWord] > wordCount[currentWord] {
				leftWord := s[left : left + wordLength]

				currentWindowCount[leftWord]--

				if currentWindowCount[leftWord] == 0 {
					delete(currentWindowCount, leftWord)
				}
				left += wordLength
			}

			if right - left == wordArrayLength*wordLength {
				result = append(result, left)
			}
		}
	}
	return result
}