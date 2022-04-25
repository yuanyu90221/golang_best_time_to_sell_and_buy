package best_sell_buy

func maxProfit(prices []int) int {
	max_profit := 0
	sell := 1
	buy := 0
	max_idx := len(prices) - 1
	for sell <= max_idx {
		profit := 0
		if prices[sell] > prices[buy] {
			profit = prices[sell] - prices[buy]
			if max_profit < profit {
				max_profit = profit
			}
		} else {
			buy = sell
		}
		sell++
	}
	return max_profit
}
