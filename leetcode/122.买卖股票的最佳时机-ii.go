package leetcode

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit122(prices []int) int {
	res := 0

	for i := 1; i < len(prices); i++ {
		res += max(prices[i]-prices[i-1], 0)
	}
	return res
}

// @lc code=end
