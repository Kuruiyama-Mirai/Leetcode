package leetcode

/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	//通过%求模取余数来模拟环
	res := make([]int, len(nums))
	stack := make([]int, 0)

	for i := 2*len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%len(nums)] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i%len(nums)] = -1
		} else {
			res[i%len(nums)] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%len(nums)])
	}
	return res
}

// @lc code=end
