package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	height := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		height[i] = envelopes[i][1]
	}
	//dp会超时 只能用2分
	return lengthOfLIS300(height)
}

func lengthOfLIS300(nums []int) int {
	top := make([]int, len(nums))
	// 牌堆数初始化为 0
	piles := 0

	for i := 0; i < len(nums); i++ {
		// 要处理的扑克牌
		poker := nums[i]

		left, right := 0, piles
		//在已有的堆里面找哪里插入，每个堆的最上层即为最大递增子序列
		for left < right {
			mid := left + (right-left)/2
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid
			}
		}
		// 没找到合适的牌堆，新建一堆
		if left == piles {
			piles++
		}
		top[left] = poker
	}
	return piles
}

// @lc code=end
