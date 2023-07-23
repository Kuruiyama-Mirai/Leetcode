package leetcode

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	res := 0

	for i := 1; i < len(prices); i++ {
		res += Max122(prices[i]-prices[i-1], 0)
	}
	return res
}

func Max122(a, b int) int {

	if a > b {
		return a
	}
	return b
}

// @lc code=end
