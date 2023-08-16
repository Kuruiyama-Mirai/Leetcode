package leetcode

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	end, far, jump := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		far = max(i+nums[i], far)
		if end == i {
			jump++
			end = far
		}
	}
	return jump
}

// @lc code=end
