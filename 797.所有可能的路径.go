package leetcode

/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	path := []int{}

	traverse797(graph, 0, &path, &res)
	return res
}
func traverse797(graph [][]int, s int, path *[]int, res *[][]int) {
	*path = append(*path, s)
	if s == len(graph)-1 {

		*res = append(*res, append([]int{}, *path...))
	}
	for _, v := range graph[s] {
		traverse797(graph, v, path, res)
	}
	*path = (*path)[:len(*path)-1]
}

// @lc code=end
