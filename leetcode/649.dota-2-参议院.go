package leetcode

/*
 * @lc app=leetcode.cn id=649 lang=golang
 *
 * [649] Dota2 参议院
 */

// @lc code=start
func predictPartyVictory(senate string) string {
	R, D := true, true
	flag := 0

	senatebyte := []byte(senate)
	for R && D {
		R = false
		D = false
		for i := 0; i < len(senatebyte); i++ {
			if senatebyte[i] == 'R' {
				if flag < 0 {
					senatebyte[i] = 0 // 消灭R，R此时为false
				} else {
					R = true // 如果没被消灭，本轮循环结束有R
				}
				flag++
			}
			if senatebyte[i] == 'D' {
				if flag > 0 {
					senatebyte[i] = 0 //消灭D
				} else {
					D = true
				}
				flag--
			}
		}
	}
	// 循环结束之后，R和D只能有一个为true
	if R {
		return "Radiant"
	}
	return "Dire"
}

// @lc code=end
