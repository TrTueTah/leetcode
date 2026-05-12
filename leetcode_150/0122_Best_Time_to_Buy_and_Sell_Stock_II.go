func maxProfit(prices []int) int {
    maxProfit := 0;
	for i := 1; i < len(prices); i++ {
		profit := max(0, prices[i] - prices[i - 1])
		maxProfit += profit
	}
	return maxProfit
}