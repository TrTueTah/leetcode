package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	minValue, max := prices[0], 0

	for _, num := range prices {
		if max < num-minValue {
			max = num - minValue
		}
		if num < minValue {
			minValue = num
		}
	}
	return max
}
