package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	isUsed := make([]bool, len(nums))

	var backtracking47 func(nums []int, isUsed []bool)
	backtracking47 = func(nums []int, isUsed []bool) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i != 0 && nums[i] == nums[i-1] && !isUsed[i-1] {
				continue
			}
			if !isUsed[i] {
				path = append(path, nums[i])
				isUsed[i] = true
				backtracking47(nums, isUsed)
				isUsed[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	//要去重剪枝一定要先排序！！！！！！
	sort.Ints(nums)
	backtracking47(nums, isUsed)
	return res
}

// @lc code=end
