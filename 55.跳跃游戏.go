package leetcode

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	cover := 0
	if len(nums) == 1 {
		return true
	}
	for i := 0; i <= cover; i++ {
		cover = max(i+nums[i], cover)

		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

// @lc code=end
