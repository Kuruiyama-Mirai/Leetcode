package leetcode

/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
func minCut(s string) int {
	//dp[i]- [0,i]范围内最少分割次数
	dp := make([]int, len(s))
	for i := range dp {
		dp[i] = i
	}

	isValid := make([][]bool, len(s))
	for i := 0; i < len(isValid); i++ {
		isValid[i] = make([]bool, len(s))
		isValid[i][i] = true
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 1 || isValid[i+1][j-1]) {
				isValid[i][j] = true
			}
		}
	}

	for i := 1; i < len(s); i++ {
		if isValid[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if isValid[j+1][i] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(s)-1]
}
func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

// @lc code=end
