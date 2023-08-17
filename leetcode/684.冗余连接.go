package leetcode

/*
 * @lc app=leetcode.cn id=684 lang=golang
 *
 * [684] 冗余连接
 */

// @lc code=start

var (
	n      = 10001
	father = make([]int, n)
)

//并查集初始化
func initalize() {
	for i := range father {
		father[i] = i
	}
}

// 并查集里寻根的过程
func find(u int) int {
	if u == father[u] {
		return u
	}
	father[u] = find(father[u])
	return father[u]
}

// 将v->u 这条边加入并查集
func join(u, v int) {
	u = find(u)
	v = find(v)
	if u == v {
		return
	}
	father[v] = u
}

// 判断 u 和 v是否找到同一个根
func same(u, v int) bool {
	u = find(u)
	v = find(v)
	return u == v
}

func findRedundantConnection(edges [][]int) []int {
	initalize()
	for i := 0; i < len(edges); i++ {
		if same(edges[i][0], edges[i][1]) {
			return edges[i]
		} else {
			join(edges[i][0], edges[i][1])
		}
	}
	return nil
}

// @lc code=end
