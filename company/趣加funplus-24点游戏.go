package company

import "math"

/*
其实就是leetcode的679题，不过改成了给定分值了
*/
func judgePoint24(cards []int, points int) bool {
	newCards := make([]float64, len(cards))
	for i := range newCards {
		newCards[i] = float64(cards[i])
	}
	return dfs(newCards, points)
}

func dfs(cards []float64, points int) bool {
	if len(cards) == 1 {
		return math.Abs(cards[0]-float64(points)) < 1e-9
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
			flag = flag || dfs(append(newNums, n1+n2), points)
			flag = flag || dfs(append(newNums, n1-n2), points)
			flag = flag || dfs(append(newNums, n2-n1), points)
			flag = flag || dfs(append(newNums, n1*n2), points)
			if n1 != 0 {
				flag = flag || dfs(append(newNums, n2/n1), points)
			}
			if n2 != 0 {
				flag = flag || dfs(append(newNums, n1/n2), points)
			}
			if flag {
				return true
			}
		}
	}
	return false
}
