package leetcode

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {

	var backtracking37 func(board [][]byte) bool
	backtracking37 = func(board [][]byte) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}
				//遍历填数字
				for k := '1'; k <= '9'; k++ {
					if isValid37(i, j, byte(k), board) {
						board[i][j] = byte(k)
						if backtracking37(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}

	backtracking37(board)
}

func isValid37(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == k {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][col] == k {
			return false
		}
	}
	//小方格
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}

// @lc code=end
