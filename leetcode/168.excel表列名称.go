package leetcode

/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	res := ""

	for columnNumber > 0 {
		mod := (columnNumber - 1) % 26
		res = string('A'+mod) + res
		columnNumber = (columnNumber - 1) / 26
	}
	return res
}

// @lc code=end
