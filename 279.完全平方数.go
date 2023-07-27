package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = Min(dp[j-i*i]+1, dp[j])
		}
	}
	return dp[n]
}

// @lc code=end
