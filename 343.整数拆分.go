package leetcode

/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1

	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = Max(dp[i], Max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

// @lc code=end
