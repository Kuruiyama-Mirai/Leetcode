package leetcode

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	res := make([]int, 26)
	for i := 0; i < len(s); i++ {
		res[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		res[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if res[i] != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
