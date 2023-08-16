package leetcode

/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	ss := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		if i+k <= len(s) {
			reverse541(ss[i : i+k])
		} else {
			reverse541(ss[i:len(s)])
		}
	}
	return string(ss)
}

func reverse541(b []byte) {
	left, right := 0, len(b)-1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

// @lc code=end
