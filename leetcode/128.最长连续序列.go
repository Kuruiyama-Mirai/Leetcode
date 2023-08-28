package leetcode

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	hash := make(map[int]bool)
	res := 0

	for _, num := range nums {
		hash[num] = true
	}
	for num := range hash {
		if hash[num-1] {
			continue
		}
		curNum := num
		curLen := 1

		for hash[curNum+1] {
			curNum += 1
			curLen += 1
		}
		res = max(res, curLen)
	}

	return res
}

// @lc code=end
