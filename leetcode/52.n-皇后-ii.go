package leetcode

/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N 皇后 II
 */

// @lc code=start
func totalNQueens(n int) int {
	res := 0
	chessboard := make([][]string, n)
	//初始化整个棋盘
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}

	var backtracking52 func(row int, chessboard [][]string)
	backtracking52 = func(row int, chessboard [][]string) {
		if row == n {
			res++
		}
		for i := 0; i < n; i++ {
			if isValid52(n, row, i, chessboard) {
				chessboard[row][i] = "Q"
				backtracking52(row+1, chessboard)
				chessboard[row][i] = "."
			}
		}
	}
	backtracking52(0, chessboard)
	return res
}

func isValid52(n, row, col int, chessboard [][]string) bool {
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	//检查左上
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	//检查右上
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}

// @lc code=end
