package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1356 lang=golang
 *
 * [1356] 根据数字二进制下 1 的数目排序
 */

// @lc code=start
func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a, b := bitCount(arr[i]), bitCount(arr[j])
		if a == b {
			return arr[i] < arr[j]
		}
		return a < b
	})
	return arr
}

func bitCount(n int) int {
	count := 0
	for n != 0 {
		n &= n - 1
		count++
	}
	return count
}

// @lc code=end
