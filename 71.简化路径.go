package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	stack := make([]string, 0)

	for _, part := range parts {
		if part == "" || part == "." {
			continue
		}
		if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, part)
		}
	}

	// 栈中存储的文件夹组成路径
	res := ""
	for i := len(stack) - 1; i >= 0; i-- {
		res = "/" + stack[i] + res
	}
	if res == "" {
		res = "/"
	}
	return res
}

// @lc code=end
