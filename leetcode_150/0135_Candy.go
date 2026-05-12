func candy(ratings []int) int {
	n := len(ratings)
	left, right := make([]int, n), make([]int, n)

	for i := range left {
		left[i] = 1
		right[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i - 1] {
			left[i] = left[i - 1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i + 1] {
			right[i] = right[i + 1] + 1
		}
	}
	total := 0
	for i := range ratings {
		total += max(left[i], right[i])
	}
	return total
}