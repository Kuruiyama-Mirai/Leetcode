package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	left, right, valid := 0, 0, 0
	start, length := 0, math.MaxInt32

	for right < len(s) {
		c := s[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++

			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

// @lc code=end
