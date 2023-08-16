package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1365 lang=golang
 *
 * [1365] 有多少小于当前数字的数字
 */

// @lc code=start
func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)

	sort.Ints(res)
	hash := make([]int, 101)
	//从后向前遍历，遇到相同的也会覆盖掉
	for i := len(res) - 1; i >= 0; i-- {
		hash[res[i]] = i
	}
	for i := range nums {
		res[i] = hash[nums[i]]
	}

	return res
}

// @lc code=end
