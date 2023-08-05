package leetcode

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	res := 0
	stack := make([]int, 0)

	//头尾要加0
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stack = append(stack, 0)

	for i := 1; i < len(heights); i++ {
		if heights[i] > heights[stack[len(stack)-1]] {
			//比栈顶大的元素入栈
			stack = append(stack, i)
		} else if heights[i] == heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && heights[i] < heights[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if len(stack) != 0 {
					temp := (i - stack[len(stack)-1] - 1) * (heights[mid])
					res = Max(res, temp)
				}
			}
			stack = append(stack, i)
		}
	}
	return res
}

// @lc code=end
