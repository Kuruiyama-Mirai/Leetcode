/*
有n辆共享单车，编号依次为A，B，C，... 。现在要将单车整齐摆放，其中A车、B车属于特殊车型，
并且B车比A车大，现要求B车必须摆在A车后，例如A-B-C， A-C-B等，有多少种摆放方法。备注：至少3辆单车。
*/

package company

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func setBikes(n int, bikes []string) (string, int) {
	res := make([]string, 0)
	ans := make([][]string, 0)
	path := make([]string, 0)
	var backtracking func(bikes []string, path []string, hasA, hasB bool)
	backtracking = func(bikes, path []string, hasA, hasB bool) {
		if len(path) == n {
			temp := make([]string, n)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := 0; i < len(bikes); i++ {
			flagA, flagB := false, false // 判断是否在当前层获取A或者B
			//保证A要在B前面
			if hasB && bikes[i] == "A" {
				continue
			} else if !hasA && bikes[i] == "B" {
				continue
			}
			if bikes[i] == "A" {
				flagA = true
				hasA = true
			} else if bikes[i] == "B" {
				flagB = true
				hasB = true
			}
			path = append(path, bikes[i])

			btmp := make([]string, 0)
			btmp = append(btmp, bikes[:i]...)
			btmp = append(btmp, bikes[i+1:]...)

			backtracking(btmp, bikes, hasA, hasB)
			path = path[:len(path)-1]
			if flagA {
				hasA = false
			}
			if flagB {
				hasB = false
			}
		}
	}
	backtracking(bikes, path, false, false)
	for _, val := range ans {
		t := strings.Join(val, "-")
		res = append(res, t)
	}

	return strings.Join(res, ""), len(res)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	bikes := strings.Split(scanner.Text(), "")
	res, resN := setBikes(n, bikes)
	fmt.Println(res, resN)
}
