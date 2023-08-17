package leetcode

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
				reverse31(nums, i+1, len(nums)-1)
				return
			}
		}
	}
	reverse31(nums, 0, len(nums)-1)
}

func reverse31(a []int, begin, end int) {
	for i, j := begin, end; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

// @lc code=end
