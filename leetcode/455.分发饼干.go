package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	child := 0
	for i := 0; child < len(g) && i < len(s); i++ {
		//用饼干跟最小的孩子胃口比
		if s[i] >= g[child] {
			child++
		}
	}
	return child
}

// @lc code=end
