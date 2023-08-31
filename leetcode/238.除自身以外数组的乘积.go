package leetcode

/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	//求前缀积
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	j := 1
	//后缀积
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * j
		j *= nums[i]
	}

	return res
}

// @lc code=end
