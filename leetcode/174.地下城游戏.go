package leetcode

/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] 地下城游戏
 */

// @lc code=start
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = max(1-dungeon[m-1][n-1], 1)
	for i := n - 2; i >= 0; i-- {
		dp[m-1][i] = max(1, dp[m-1][i+1]-dungeon[m-1][i])
	}
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = max(1, dp[i+1][n-1]-dungeon[i][n-1])
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = min(max(1, dp[i][j+1]-dungeon[i][j]), max(1, dp[i+1][j]-dungeon[i][j]))
		}
	}
	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
