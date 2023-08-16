package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=1002 lang=golang
 *
 * [1002] 查找共用字符
 */

// @lc code=start
func commonChars(words []string) []string {
	res := make([]string, 0)
	minfreq := make([]int, 26)
	for i := range minfreq {
		minfreq[i] = math.MaxInt64
	}

	for _, word := range words {
		freq := [26]int{}
		for _, s := range word {
			freq[s-'a']++
		}
		for i, f := range freq {
			if f < minfreq[i] {
				minfreq[i] = f
			}
		}
	}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < minfreq[i]; j++ {
			res = append(res, string('a'+i))
		}
	}
	return res
}

// @lc code=end
