package company

import "sort"

/*
其实就是leetcode的剑指offer61题，不过改成了A可以变成14
*/

func isStraight(nums []int) bool {
	zero := 0
	for i := range nums {
		if nums[i] == 0 {
			zero++
		}
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != 0 && nums[i] == nums[i+1] {
			return false
		}
	}
	if nums[len(nums)-1]-nums[zero] > 5 {
		for i := range nums {
			if nums[i] == 1 {
				nums[i] = 14
			}
		}
		sort.Ints(nums)
		return nums[len(nums)-1]-nums[zero] < 5
	}
	return nums[len(nums)-1]-nums[zero] < 5
}
