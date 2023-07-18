package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSumVer1(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		ans := target - nums[i]
		if _, ok := m[ans]; ok {
			return []int{m[ans], i}
		}
		m[nums[i]] = i
	}
	return nil
}

// @lc code=end
