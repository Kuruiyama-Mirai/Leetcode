package leetcode

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	i, sum, j := 0, 0, 0
	res := len(nums) + 1
	for j < len(nums) {
		sum += nums[j]

		for sum >= target {
			sublength := j - i + 1
			if sublength < res {
				res = sublength
			}
			sum -= nums[i]

			i++
		}
		j++
	}
	if res == len(nums)+1 {
		return 0
	} else {
		return res
	}
}

// @lc code=end
