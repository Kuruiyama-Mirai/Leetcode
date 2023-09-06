/*
自从学习了计算机编程，小明就彻底 爱上了0和1，很多01串经第在他脑子里面浮现。
但是他从小就害怕警察，虽然他并没有犯过啥大错误，最多就是在小花的新衣服上画朵小花，在小红的白裙子上点上几个小红点…所以他很害怕110。现在给你一个01串，请你编写一个程序计算在这个01串中不包含110的最长子串的长度。输入描述：
输入一行，包含一个01串S（在S中只包含0和1，有可能输入全为0或者1的
串) , 其长度1≤|S|≤1000000。输出描述：输出01串S中不包含110的最长子串的长度。

*/

package company

import (
	"strings"
)

func test(S string) int {
	if strings.Count(S, "0") == len(S) || strings.Count(S, "1") == len(S) {
		return len(S)
	}
	if len(S) <= 3 {
		return len(S)
	}
	res := 0
	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {
			if strings.Contains(S[i:j], "110") {
				break
			}
			temp := j - i
			res = max(temp, res)
		}
	}

	return res
}
