package leetcode

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := make([]string, 0)

	var backtracking22 func(l, r int, path string)
	backtracking22 = func(l, r int, path string) {
		if len(path) == 2*n {
			res = append(res, path)
		}
		if l > 0 {
			backtracking22(l-1, r, path+"(")
		}
		if l < r {
			backtracking22(l, r-1, path+")")
		}
	}
	backtracking22(n, n, "")
	return res
}

// @lc code=end
