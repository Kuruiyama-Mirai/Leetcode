package leetcode

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob198(nums []int) int {
	dp := make([]int, len(nums)+1)
	if len(nums) == 1 {
		return nums[0]
	}
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = Max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

// @lc code=end
