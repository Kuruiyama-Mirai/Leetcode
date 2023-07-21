package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtracking90 func(nums []int, startIndex int)
	backtracking90 = func(nums []int, startIndex int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		for i := startIndex; i < len(nums); i++ {
			if i != startIndex && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtracking90(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	sort.Ints(nums)
	backtracking90(nums, 0)
	return res
}

// @lc code=end
