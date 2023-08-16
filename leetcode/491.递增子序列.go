package leetcode

/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtracking491 func(nums []int, startIndex int)
	backtracking491 = func(nums []int, startIndex int) {
		if len(path) >= 2 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		used := make(map[int]bool, len(nums)) //对同层元素去重
		for i := startIndex; i < len(nums); i++ {
			if used[nums[i]] { //去重
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				path = append(path, nums[i])
				used[nums[i]] = true
				backtracking491(nums, i+1)
				path = path[:len(path)-1]
			}
		}
	}
	backtracking491(nums, 0)
	return res
}

// @lc code=end
