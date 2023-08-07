package leetcode

/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	count := make([]int, 2002)
	for i := 0; i < len(arr); i++ {
		count[arr[i]+1000]++
	}

	frequence := make([]bool, 1002)
	for i := 0; i <= 2000; i++ {
		if count[i] > 0 {
			if !frequence[count[i]] {
				frequence[count[i]] = true
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=end
