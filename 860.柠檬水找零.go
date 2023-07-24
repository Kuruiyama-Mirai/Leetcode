package leetcode

/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for i := range bills {
		if bills[i] == 5 {
			five++
		}
		if bills[i] == 10 {
			if five < 0 {
				return false
			}
			ten++
			five--
		}
		if bills[i] == 20 {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=end
