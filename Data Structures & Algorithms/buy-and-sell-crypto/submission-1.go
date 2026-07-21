func maxProfit(prices []int) int {
	profit := 0
	n := len(prices)

	buy_price := prices[0]
	for i := 0; i < n; i++ {
		buy_price = min(prices[i], buy_price)
		profit = max(profit, prices[i] - buy_price) 
	}

	return profit
}
