package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	isUsed := make([]bool, len(candidates))

	var backtracking40 func(candidates []int, target, sum, startIndex int, isUsed []bool)
	backtracking40 = func(candidates []int, target, sum, startIndex int, isUsed []bool) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := startIndex; i < len(candidates); i++ {
			// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
			// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
			if i > 0 && candidates[i] == candidates[i-1] && !isUsed[i-1] {
				continue
			}
			sum += candidates[i]
			isUsed[i] = true
			path = append(path, candidates[i])
			backtracking40(candidates, target, sum, i+1, isUsed)
			sum -= candidates[i]
			isUsed[i] = false
			path = path[:len(path)-1]
		}
	}
	sort.Ints(candidates)
	backtracking40(candidates, target, 0, 0, isUsed)
	return res
}

// @lc code=end
