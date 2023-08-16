package leetcode

/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 */

// @lc code=start
func removeDuplicateLetters(s string) string {
	count := [26]int{}
	for _, c := range s {
		count[c-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if count[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		count[ch-'a']--
	}
	return string(stack)
}

// @lc code=end
