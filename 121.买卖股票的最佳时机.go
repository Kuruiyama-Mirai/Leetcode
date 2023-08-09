package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit121(prices []int) int {
	low := math.MaxInt32
	res := 0
	for i := 0; i < len(prices); i++ {
		low = min(low, prices[i])
		res = max(res, prices[i]-low)
	}
	return res
}

// @lc code=end
