package leetcode

/*
 * @lc app=leetcode.cn id=925 lang=golang
 *
 * [925] 长按键入
 */

// @lc code=start
func isLongPressedName(name string, typed string) bool {
	if name[0] != typed[0] || len(name) > len(typed) {
		return false
	}
	index := 0
	var lastSame byte

	for i := 0; i < len(typed); i++ {
		if index < len(name) && name[index] == typed[i] {
			lastSame = name[index]
			index++
		} else if typed[i] == lastSame {
			continue
		} else {
			return false
		}
	}
	return index == len(name)
}

// @lc code=end
