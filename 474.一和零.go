package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		zeroNum := strings.Count(strs[i], "0")
		oneNum := len(strs[i]) - zeroNum
		//m,n就是背包容量 限制死0，1的数量
		for j := m; j >= zeroNum; j-- {
			for k := n; k >= oneNum; k-- {
				dp[j][k] = Max(dp[j][k], dp[j-zeroNum][k-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}

// @lc code=end
