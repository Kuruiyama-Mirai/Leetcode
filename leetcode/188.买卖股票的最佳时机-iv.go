package leetcode

/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit188(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2*k+1)
	}
	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		//j 奇数买入 偶数卖出
		for j := 0; j < 2*k-1; j += 2 {
			dp[i][j+1] = max(dp[i-1][j]-prices[i], dp[i-1][j+1])
			dp[i][j+2] = max(dp[i-1][j+1]+prices[i], dp[i-1][j+2])
		}
	}
	return dp[len(prices)-1][2*k]
}

// @lc code=end
