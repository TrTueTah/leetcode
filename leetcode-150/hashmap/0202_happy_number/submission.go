package happynumber

func isHappy(n int) bool {
	store := make(map[int]bool)

	for !store[n] {
		store[n] = true
		n = getNextNumber(n)
		if n == 1 {
			return true
		}
	}

	return false
}

func getNextNumber(n int) int {
	new := 0
	for n > 0 {
		digit := n % 10
		new += digit * digit
		n /= 10
	}
	return new
}
