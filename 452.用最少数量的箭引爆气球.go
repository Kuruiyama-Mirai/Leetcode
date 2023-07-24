package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	res := 1
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	for i := 1; i < len(points); i++ {
		if points[i-1][1] < points[i][0] {
			res++
		} else {
			points[i][1] = Min452(points[i][1], points[i-1][1])
		}
	}
	return res
}

func Min452(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
