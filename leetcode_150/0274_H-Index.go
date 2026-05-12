func hIndex(citations []int) int {
	sort.Ints(citations)

	k := 0
	n := len(citations)

	for k < n && citations[n-1-k] > k {
		k++
	}

	return k
}