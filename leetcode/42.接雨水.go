package leetcode

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	res := 0
	//单调栈只用存下标就行
	stack := make([]int, 1)

	for i := 1; i < len(height); i++ {
		if height[i] < height[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if height[i] == height[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if len(stack) != 0 {
					water := (min(height[i], height[stack[len(stack)-1]]) - height[top]) * (i - stack[len(stack)-1] - 1)
					res += water
				}
			}
			stack = append(stack, i)
		}
	}
	return res
}

// @lc code=end
