package isomorphicstrings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	store1, store2 := make(map[byte]byte), make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		c1 := s[i]
		c2 := t[i]

		if val, exist := store1[c1]; exist && val != c2 {
			return false
		}

		if val, exist := store2[c2]; exist && val != c1 {
			return false
		}

		store1[c1] = c2
		store2[c2] = c1
	}

	return true
}
