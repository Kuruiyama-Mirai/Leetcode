package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	nums := make([]int, len(s))
	for i := 0; i < len(nums); i++ {
		switch s[i] {
		case 'A':
			nums[i] = 0
		case 'G':
			nums[i] = 1
		case 'C':
			nums[i] = 2
		case 'T':
			nums[i] = 3
		}
	}
	//记录重复出现的哈希值
	seen := make(map[int]bool)
	res := make(map[string]bool)

	var sequences []string

	// 数字位数
	L := 10
	// 进制
	R := 4
	// 存储 R^(L - 1) 的结果
	RL := int(math.Pow(float64(R), float64(L-1)))
	// 维护滑动窗口中字符串的哈希值
	windowHash := 0

	left, right := 0, 0
	for right < len(nums) {
		windowHash = R*windowHash + nums[right]
		right++

		if right-left == L {
			if seen[windowHash] {
				res[s[left:right]] = true
			} else {
				seen[windowHash] = true
			}
			windowHash = windowHash - RL*nums[left]
			left++
		}
	}
	for k := range res {
		sequences = append(sequences, k)
	}
	return sequences
}

// @lc code=end
