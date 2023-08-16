package leetcode

/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtracking216 func(k, n, sum, startIndex int)
	backtracking216 = func(k, n, sum, startIndex int) {
		if sum > n {
			return
		}
		if len(path) == k {
			if sum == n {
				temp := make([]int, k)
				copy(temp, path)
				res = append(res, temp)
				return
			}
		}
		for i := startIndex; i <= 9-(k-len(path))+1; i++ {
			sum += i
			path = append(path, i)
			backtracking216(k, n, sum, i+1)
			sum -= i
			path = path[:len(path)-1]
		}
	}

	backtracking216(k, n, 0, 1)
	return res
}

// @lc code=end
