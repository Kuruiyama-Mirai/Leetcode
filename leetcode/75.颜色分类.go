package leetcode

/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	index, left, right := 0, 0, len(nums)-1
	for index <= right {
		if nums[index] == 0 {
			nums[left], nums[index] = nums[index], nums[left]
			left++
			index++
		} else if nums[index] == 1 {
			index++
		} else if nums[index] == 2 {
			nums[right], nums[index] = nums[index], nums[right]
			right--
		}
	}
}

// @lc code=end
