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
		cover = Max55(i+nums[i], cover)

		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

func Max55(a, b int) int {

	if a > b {
		return a
	}
	return b
}

// @lc code=end
