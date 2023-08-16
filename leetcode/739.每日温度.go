package leetcode

/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	// 这里放元素索引，而不是元素
	s := make([]int, 0)

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(s) > 0 && temperatures[s[len(s)-1]] <= temperatures[i] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			res[i] = 0
		} else {
			res[i] = s[len(s)-1] - i
		}
		s = append(s, i)
	}
	return res
}

// @lc code=end
