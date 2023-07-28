package leetcode

/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob213(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return Max(nums[0], nums[1])
	}
	result1 := robRange(nums, 0, len(nums)-2)
	result2 := robRange(nums, 1, len(nums)-1)
	return Max(result1, result2)
}

//选择偷取的范围
func robRange(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	dp := make([]int, len(nums)+1)
	dp[start] = nums[start]
	dp[start+1] = Max(nums[start], nums[start+1])

	for i := start + 2; i <= end; i++ {
		dp[i] = Max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[end]
}

// @lc code=end
