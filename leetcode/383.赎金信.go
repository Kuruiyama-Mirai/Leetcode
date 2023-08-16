package leetcode

/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	res := make([]int, 26)
	for _, v := range magazine {
		res[v-'a']++
	}

	for _, v := range ransomNote {
		res[v-'a']--
		if res[v-'a'] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end
