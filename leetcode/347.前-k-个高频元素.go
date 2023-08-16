package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0)
	map_num := map[int]int{}

	for _, num := range nums {
		map_num[num]++
	}

	for key, _ := range map_num {
		res = append(res, key)
	}

	sort.Slice(res, func(i, j int) bool {
		return map_num[res[i]] > map_num[res[j]]
	})
	return res[:k]
}

// @lc code=end
