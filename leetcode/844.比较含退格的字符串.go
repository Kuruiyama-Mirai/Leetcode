package leetcode

/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
func backspaceCompare(s string, t string) bool {
	return getString(s) == getString(t)
}

func getString(s string) string {
	stack := []rune{}
	for _, c := range s {
		if c != '#' {
			stack = append(stack, c)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

// @lc code=end
