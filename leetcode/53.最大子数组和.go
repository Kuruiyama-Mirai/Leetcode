package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	res := math.MinInt32
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > res {
			res = count
		}
		if count <= 0 {
			count = 0
		}
	}
	return res
}

// @lc code=end
