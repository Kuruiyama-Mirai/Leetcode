package leetcode

/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心下标
 */

// @lc code=start
func pivotIndex(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	left, right := 0, 0
	for i := 0; i < len(nums); i++ {
		left += nums[i]
		right = sum - left + nums[i]

		if right == left {
			return i
		}
	}
	return -1
}

// @lc code=end
