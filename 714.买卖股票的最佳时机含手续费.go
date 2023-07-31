package leetcode

/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit714(prices []int, fee int) int {
	dp := make([][2]int, len(prices))

	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}

	return Max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

// @lc code=end
