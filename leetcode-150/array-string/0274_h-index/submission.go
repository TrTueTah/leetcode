package hindex

import "slices"

func hIndex(citations []int) int {
	slices.Sort(citations)

	for i := len(citations) - 1; i >= 0; i-- {
		if len(citations)-i > citations[i] {
			return len(citations) - i - 1
		}
	}

	return len(citations)
}
