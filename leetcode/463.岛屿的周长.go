package leetcode

/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 */

// @lc code=start
func islandPerimeter(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				//判断上边界
				if i-1 < 0 || grid[i-1][j] == 0 {
					res++
				}
				//判断下边界
				if i+1 >= len(grid) || grid[i+1][j] == 0 {
					res++
				}
				//判断左边界
				if j-1 < 0 || grid[i][j-1] == 0 {
					res++
				}
				//判断右边界
				if j+1 >= len(grid[0]) || grid[i][j+1] == 0 {
					res++
				}
			}
		}
	}
	return res
}

// @lc code=end
