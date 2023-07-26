package leetcode

/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	dp := make([]int, 10001)
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = Max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}

// @lc code=end
