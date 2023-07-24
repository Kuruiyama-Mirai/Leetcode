package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 */

// @lc code=start
func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	ss := []byte(s)

	if len(ss) <= 1 {
		return n
	}
	for i := len(ss) - 1; i > 0; i-- {
		if ss[i-1] > ss[i] {
			ss[i-1] -= 1
			for j := i; j < len(ss); j++ {
				ss[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(ss))

	return res
}

// @lc code=end
