package leetcode

/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 记录 nums2 中每个元素的下一个更大元素
	greater := nextGreater(nums2)
	// 转化成映射：元素 x -> x 的下一个最大元素
	greaterMap := make(map[int]int)
	for i, v := range nums2 {
		greaterMap[v] = greater[i]
	}
	// nums1 是 nums2 的子集，所以根据 greaterMap 可以得到结果
	res := make([]int, len(nums1))
	for i, v := range nums1 {
		res[i] = greaterMap[v]
	}
	return res
}

func nextGreater(nums []int) []int {
	res := make([]int, len(nums))
	stack := make([]int, 0)

	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return res
}

// @lc code=end
