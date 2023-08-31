package leetcode

/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = countOne(i)
	}
	return res
}

func countOne(n int) int {
	res := 0
	for ; n > 0; n &= n - 1 {
		res++
	}
	return res
}

// @lc code=end
