package leetcode

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition131(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	var backtracking131 func(s string, startIndex int)
	backtracking131 = func(s string, startIndex int) {
		if startIndex >= len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := startIndex; i < len(s); i++ {
			str := s[startIndex : i+1]

			if isPalindrome131(str, 0, len(str)-1) {
				path = append(path, str)
				backtracking131(s, i+1)
				path = path[:len(path)-1]
			}
		}
	}
	backtracking131(s, 0)
	return res
}

func isPalindrome131(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// @lc code=end
