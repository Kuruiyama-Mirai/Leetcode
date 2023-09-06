package leetcode

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak139(s string, wordDict []string) bool {
	words := make(map[string]bool)

	for _, w := range wordDict {
		words[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end
