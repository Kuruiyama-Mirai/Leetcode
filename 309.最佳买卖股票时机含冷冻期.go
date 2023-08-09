package leetcode

/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit309(prices []int) int {

	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		//0 今天买入或一直有 1 保持卖出 2 今天卖出 3 冷冻
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3], dp[i-1][1])-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(dp[len(prices)-1][3], max(dp[len(prices)-1][1], dp[len(prices)-1][2]))
}

// @lc code=end
