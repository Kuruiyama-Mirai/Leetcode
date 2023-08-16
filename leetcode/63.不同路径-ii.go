package leetcode

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	//先排除特殊情况
	if obstacleGrid[row-1][col-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	//行列都没有障碍物时才可以初始化为1
	for i := 0; i < row && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for j := 1; j < col && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[row-1][col-1]
}

// @lc code=end
