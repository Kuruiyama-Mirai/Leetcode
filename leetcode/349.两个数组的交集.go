package leetcode

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	m := map[int]bool{}

	for _, num := range nums1 {
		m[num] = true
	}
	for _, num := range nums2 {
		if m[num] {
			m[num] = false
			res = append(res, num)
		}
	}

	return res
}

// @lc code=end
