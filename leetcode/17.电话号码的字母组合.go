package leetcode

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	letterMap := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := make([]string, 0)
	path := make([]byte, 0)

	if digits == "" {
		return res
	}
	var backtracking17 func(digits string, index int)
	backtracking17 = func(digits string, index int) {
		if len(path) == len(digits) {
			temp := string(path)
			res = append(res, temp)
			return
		}
		digit := int(digits[index] - '0')
		letters := letterMap[digit]
		for i := 0; i < len(letters); i++ {
			path = append(path, []byte(letters)[i])
			backtracking17(digits, index+1)
			path = path[:len(path)-1]
		}
	}

	backtracking17(digits, 0)
	return res
}

// @lc code=end
