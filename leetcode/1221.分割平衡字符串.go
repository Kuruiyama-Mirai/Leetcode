package leetcode

/*
 * @lc app=leetcode.cn id=1221 lang=golang
 *
 * [1221] 分割平衡字符串
 */

// @lc code=start
func balancedStringSplit(s string) int {
	res, count := 0, 0
	for i := range s {
		if s[i] == 'R' {
			count++
		} else {
			count--
		}
		if count == 0 {
			res++
		}
	}
	return res
}

// @lc code=end
