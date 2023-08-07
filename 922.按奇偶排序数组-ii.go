package leetcode

/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 */

// @lc code=start
func sortArrayByParityII(nums []int) []int {
	oldIndex := 1

	for i := 0; i < len(nums); i += 2 {
		//如果在偶数位遇到奇数
		if nums[i]%2 == 1 {
			for nums[oldIndex]%2 != 0 {
				//在奇数位中找一个偶数
				oldIndex += 2
			}
			swap922(nums, i, oldIndex)
		}
	}
	return nums
}

func swap922(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

// @lc code=end
