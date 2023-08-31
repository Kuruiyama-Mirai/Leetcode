package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=931 lang=golang
 *
 * [931] 下降路径最小和
 */

// @lc code=start
func minFallingPathSum(matrix [][]int) int {
	res := math.MaxInt64
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix[0]); i++ {
		dp[0][i] = matrix[0][i]
	}
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if j < 1 {
				dp[i][j] = matrix[i][j] + min(dp[i-1][j], dp[i-1][j+1])
			} else if j >= len(matrix[0])-1 {
				dp[i][j] = matrix[i][j] + min(dp[i-1][j], dp[i-1][j-1])
			} else {
				dp[i][j] = matrix[i][j] + min(min(dp[i-1][j-1], dp[i-1][j+1]), dp[i-1][j])
			}
		}

		for i := 0; i < len(matrix); i++ {
			res = min(res, dp[len(matrix)-1][i])
		}
	}
	return res
}

// @lc code=end
