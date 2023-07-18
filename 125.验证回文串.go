package leetcode

import (
	"strings"
	"unicode"
)

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		cs := s[i]
		if unicode.IsLetter(rune(cs)) || unicode.IsDigit(rune(cs)) {
			sb.WriteRune(unicode.ToLower(rune(cs)))
		}
	}

	s = sb.String()
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// @lc code=end
