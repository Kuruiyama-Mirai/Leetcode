package leetcode

/*
 * @lc app=leetcode.cn id=841 lang=golang
 *
 * [841] 钥匙和房间
 */

// @lc code=start
func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))

	var dfs func(key int, rooms [][]int, visited []bool)
	dfs = func(key int, rooms [][]int, visited []bool) {
		if visited[key] {
			return
		}
		visited[key] = true
		keys := rooms[key]
		for _, key := range keys {
			dfs(key, rooms, visited)
		}
	}

	dfs(0, rooms, visited)
	for i := range visited {
		if !visited[i] {
			return false
		}
	}
	return true
}

// @lc code=end
