package leetcode

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	//初始化二维数组
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	up, down, left, right := 0, n-1, 0, n-1
	num := 1

	for num <= n*n {
		if up <= down {
			for i := left; i <= right; i++ {
				res[up][i] = num
				num++
			}
			up++
		}
		if left <= right {
			for i := up; i <= down; i++ {
				res[i][right] = num
				num++
			}
			right--
		}
		if up <= down {
			for i := right; i >= left; i-- {
				res[down][i] = num
				num++
			}
			down--
		}
		if left <= right {
			for i := down; i >= up; i-- {
				res[i][left] = num
				num++
			}
			left++
		}
	}
	return res
}

// @lc code=end
