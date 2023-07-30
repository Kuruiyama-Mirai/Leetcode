package leetcode

/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit123(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 5)
	}
	//dp[i][j]表示第i天j状态所持有的钱 1买 2卖 3再买 4再卖
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = Max(dp[i-1][0]-prices[i], dp[i-1][1])
		dp[i][2] = Max(dp[i-1][1]+prices[i], dp[i-1][2])
		dp[i][3] = Max(dp[i-1][2]-prices[i], dp[i-1][3])
		dp[i][4] = Max(dp[i-1][3]+prices[i], dp[i-1][4])
	}
	return dp[len(prices)-1][4]
}

// @lc code=end
