package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if int(math.Abs(float64(target))) > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	bag := (sum + target) / 2

	dp := make([]int, bag+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bag]
}

// @lc code=end
