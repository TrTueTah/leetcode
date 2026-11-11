package validanagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	store := make(map[rune]int)

	for _, ch := range s {
		store[ch]++
	}

	for _, ch := range t {
		if store[ch] == 0 {
			return false
		}
		store[ch]--
	}

	return true
}
