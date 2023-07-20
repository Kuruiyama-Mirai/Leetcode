package leetcode

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtracking39 func(candidates []int, target, sum, startIndex int)
	backtracking39 = func(candidates []int, target, sum, startIndex int) {
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
			sum += candidates[i]
			path = append(path, candidates[i])
			backtracking39(candidates, target, sum, i) //index不用i+1,因为元素可以重复选取
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	backtracking39(candidates, target, 0, 0)
	return res
}

// @lc code=end
