package leetcode

/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			//j-1 为j为头结点左子树节点数量，i-j 为以j为头结点右子树节点数量
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

// @lc code=end
