package array

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/description/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			sum += prices[i+1] - prices[i]
		}
	}
	return sum
}
