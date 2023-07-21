package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
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
	var backtracking51 func(row int, chessboard [][]string)
	backtracking51 = func(row int, chessboard [][]string) {
		if row == n {
			temp := make([]string, n)
			for i, rorstr := range chessboard {
				temp[i] = strings.Join(rorstr, "")
			}
			res = append(res, temp)
		}
		for i := 0; i < n; i++ {
			if isValid51(n, row, i, chessboard) {
				chessboard[row][i] = "Q"
				backtracking51(row+1, chessboard)
				chessboard[row][i] = "."
			}
		}
	}
	backtracking51(0, chessboard)
	return res
}

func isValid51(n, row, col int, chessboard [][]string) bool {
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
