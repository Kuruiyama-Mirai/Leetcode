package leetcode

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	res := make([]int, 0, len(matrix)*len(matrix[0]))

	for len(res) < len(matrix)*len(matrix[0]) {
		if up <= down {
			for i := left; i <= right; i++ {
				res = append(res, matrix[up][i])
			}
			up++
		}
		if left <= right {
			for i := up; i <= down; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}
		if up <= down {
			for i := right; i >= left; i-- {
				res = append(res, matrix[down][i])
			}
			down--
		}
		if left <= right {
			for i := down; i >= up; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}

// @lc code=end
