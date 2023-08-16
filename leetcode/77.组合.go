package leetcode

/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtracking77 func(n, k, startIndex int)
	backtracking77 = func(n, k, startIndex int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := startIndex; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtracking77(n, k, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking77(n, k, 1)
	return res
}

// @lc code=end
