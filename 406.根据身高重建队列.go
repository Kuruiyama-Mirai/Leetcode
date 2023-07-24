package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			//身高相同按k排序
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	//再按k排序，优先插入k小的
	for _, p := range people {
		res = append(res, p)
		copy(res[p[1]+1:], res[p[1]:])
		res[p[1]] = p
	}
	return res
}

// @lc code=end
