package romantointeger

func romanToInt(s string) int {
	value := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0

	for i := 0; i < len(s)-1; i++ {
		if value[s[i]] < value[s[i+1]] {
			result -= value[s[i]]
		} else {
			result += value[s[i]]
		}
	}

	result += value[s[len(s)-1]]
	return result
}
