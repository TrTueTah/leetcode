package wordpattern

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)

	if len(words) != len(pattern) {
		return false
	}

	store1, store2 := make(map[byte]string), make(map[string]byte)

	for i := 0; i < len(words); i++ {
		ch := pattern[i]
		word := words[i]

		if val, exist := store1[ch]; exist && val != word {
			return false
		}

		if val, exist := store2[word]; exist && val != ch {
			return false
		}
		store1[ch] = word
		store2[word] = ch
	}

	return true
}
