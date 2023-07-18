package leetcode

import "container/list"

/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *ChildrenNode) [][]int {
	queue := list.New()
	res := [][]int{}

	if root == nil {
		return res
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		var temp []int
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*ChildrenNode)
			temp = append(temp, node.Val)
			for i := 0; i < len(node.Children); i++ {
				queue.PushBack(node.Children[i])
			}
		}
		res = append(res, temp)
	}
	return res
}

// @lc code=end
