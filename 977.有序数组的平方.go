package leetcode

/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	i, j, k := 0, len(nums)-1, len(nums)-1
	for i <= j {
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]
		if lm > rm {
			res[k] = lm
			i++
		} else {
			res[k] = rm
			j--
		}
		k--
	}

	return res
}

// @lc code=end
