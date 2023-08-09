package leetcode

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) <= 1 {
		return 1
	}
	for i := range dp {
		dp[i] = 1
	}
	res := 0

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// @lc code=end
