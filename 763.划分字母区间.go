package leetcode

/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
func partitionLabels(s string) []int {
	res := make([]int, 0)
	letter := make([]int, 26)

	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		letter[s[i]-'a'] = i
	}
	for i := 0; i < len(s); i++ {
		right = Max763(right, letter[s[i]-'a'])
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}
	return res
}

func Max763(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
