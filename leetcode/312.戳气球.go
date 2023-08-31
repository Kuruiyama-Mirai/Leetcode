package leetcode

/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

// @lc code=start
func maxCoins(nums []int) int {
	points := make([]int, len(nums)+2)
	points[0] = 1
	points[len(nums)+1] = 1
	for i := 1; i <= len(nums); i++ {
		points[i] = nums[i-1]
	}
	dp := make([][]int, len(nums)+2)
	for i := range dp {
		dp[i] = make([]int, len(nums)+2)
	}
	for i := len(nums); i >= 0; i-- {
		for j := i + 1; j < len(nums)+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[k]*points[j])
			}
		}
	}

	return dp[0][len(nums)+1]
}

// @lc code=end
