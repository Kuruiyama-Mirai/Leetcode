package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	path := make([]string, 0, len(s))
	var backtracking93 func(s string, startIndex int)
	backtracking93 = func(s string, startIndex int) {
		if len(path) == 4 {
			if startIndex == len(s) {
				str := strings.Join(path, ".")
				res = append(res, str)
			}
			return
		}
		for i := startIndex; i < len(s); i++ {
			if i != startIndex && s[startIndex] == '0' {
				continue
			}
			str := s[startIndex : i+1]
			num, _ := strconv.Atoi(str)
			if num >= 0 && num <= 255 {
				path = append(path, str)
				backtracking93(s, i+1)
				path = path[:len(path)-1]
			} else {
				break
			}
		}
	}

	backtracking93(s, 0)
	return res
}

// @lc code=end
