package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=679 lang=golang
 *
 * [679] 24 点游戏
 */

// @lc code=start
func judgePoint24(cards []int) bool {
	newCards := make([]float64, len(cards))
	for i := range newCards {
		newCards[i] = float64(cards[i])
	}
	return dfs679(newCards)
}

func dfs679(cards []float64) bool {
	if len(cards) == 1 {
		return math.Abs(cards[0]-24) < 1e-9
	}
	flag := false
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			n1, n2 := cards[i], cards[j]
			newNums := make([]float64, 0)
			for k := 0; k < len(cards); k++ {
				if k != i && k != j {
					newNums = append(newNums, cards[k])
				}
			}
			flag = flag || dfs679(append(newNums, n1+n2))
			flag = flag || dfs679(append(newNums, n1-n2))
			flag = flag || dfs679(append(newNums, n2-n1))
			flag = flag || dfs679(append(newNums, n1*n2))
			if n1 != 0 {
				flag = flag || dfs679(append(newNums, n2/n1))
			}
			if n2 != 0 {
				flag = flag || dfs679(append(newNums, n1/n2))
			}
			if flag {
				return true
			}
		}
	}
	return false
}

// @lc code=end
