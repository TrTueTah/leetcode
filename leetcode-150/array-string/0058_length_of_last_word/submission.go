package lengthoflastword

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	result := 0
	for i >= 0 && s[i] == ' ' {
		i--
	}
	for i >= 0 && s[i] != ' ' {
		i--
		result++
	}
	return result
}
