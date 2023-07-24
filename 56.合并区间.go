package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if right < intervals[i][0] {
			res = append(res, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else {
			right = Max56(right, intervals[i][1])
		}
	}
	//要处理最后一个区间
	res = append(res, []int{left, right})
	return res
}

func Max56(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
