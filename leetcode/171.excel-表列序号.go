package leetcode

/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel 表列序号
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
	var res int

	for i := 0; i < len(columnTitle); i++ {
		val := int(columnTitle[i] - 'A' + 1)
		res = res*26 + val
	}

	return res
}

// @lc code=end
