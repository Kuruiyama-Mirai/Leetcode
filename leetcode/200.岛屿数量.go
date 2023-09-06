package leetcode

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	res := 0
	var dfs200 func(grid [][]byte, i, j int)
	dfs200 = func(grid [][]byte, i, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs200(grid, i+1, j)
		dfs200(grid, i, j+1)
		dfs200(grid, i-1, j)
		dfs200(grid, i, j-1)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs200(grid, i, j)
			}
		}
	}
	return res
}

// @lc code=end
