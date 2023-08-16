package leetcode

/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// @lc code=end
