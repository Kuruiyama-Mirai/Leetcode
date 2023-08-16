package leetcode

/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
 */

// @lc code=start
func findNumberOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	//dp[i] i之前的最长递增子序列长度
	dp := make([]int, len(nums))
	//以nums[i]结尾的 最长递增子序列的个数
	count := make([]int, len(nums))
	maxCount, res := 0, 0
	for i := range dp {
		dp[i] = 1
		count[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					count[i] = count[j]
					dp[i] = dp[j] + 1
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
		if dp[i] > maxCount {
			maxCount = dp[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if maxCount == dp[i] {
			res += count[i]
		}
	}
	return res
}

// @lc code=end
