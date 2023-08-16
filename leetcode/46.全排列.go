package leetcode

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	isUsed := make([]bool, len(nums))

	var backtracking46 func(nums []int, isUsed []bool)
	backtracking46 = func(nums []int, isUsed []bool) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !isUsed[i] {
				path = append(path, nums[i])
				isUsed[i] = true
				backtracking46(nums, isUsed)
				isUsed[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	backtracking46(nums, isUsed)

	return res
}

// @lc code=end
