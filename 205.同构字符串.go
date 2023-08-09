package leetcode

/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	map1 := make(map[byte]byte)
	map2 := make(map[byte]byte)

	for i := range s {
		if _, ok := map1[s[i]]; !ok {
			map1[s[i]] = t[i]
		}
		if _, ok := map2[t[i]]; !ok {
			map2[t[i]] = s[i]
		}
		if map1[s[i]] != t[i] || map2[t[i]] != s[i] {
			return false
		}
	}
	return true
}

// @lc code=end
