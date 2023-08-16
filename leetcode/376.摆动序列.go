package leetcode

/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	res := 1
	prediff := nums[1] - nums[0]
	if prediff != 0 {
		res = 2
	}
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prediff <= 0 || diff < 0 && prediff >= 0 {
			res++
			prediff = diff
		}
	}
	return res
}

// @lc code=end
