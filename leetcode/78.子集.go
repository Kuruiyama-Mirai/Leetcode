package leetcode

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtracking78 func(nums []int, startIndex int)
	backtracking78 = func(nums []int, startIndex int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp) //收集子集要放在最上面， 否则会漏掉自己

		if startIndex >= len(nums) {
			return
		}

		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking78(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	backtracking78(nums, 0)
	return res
}

// @lc code=end
