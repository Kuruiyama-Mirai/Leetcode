package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	words := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		words[v] = true
	}
	//先通过DP判断字符串是否能被拆分
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	res := []string{}
	if !dp[len(s)] {
		return res
	}
	//回溯走起
	var backtracking140 func(ns string, path []string)
	backtracking140 = func(ns string, path []string) {
		if len(ns) == 0 {
			res = append(res, strings.Join(path, " "))
			return
		}
		for i := 1; i <= len(ns); i++ {
			if words[ns[:i]] {
				backtracking140(ns[i:], append(path, ns[:i]))
			}
		}
	}
	backtracking140(s, []string{})
	return res
}

// @lc code=end
