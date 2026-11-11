package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	store := make(map[rune]int)

	for _, ch := range magazine {
		store[ch]++
	}

	for _, ch := range ransomNote {
		if store[ch] == 0 {
			return false
		}
		store[ch]--
	}

	return true
}
