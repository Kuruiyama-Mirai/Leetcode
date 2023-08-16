package leetcode

/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	sum := 0
	for len(q) > 0 {
		sz := len(q)
		level := []int{}

		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			sum += cur.Val
		}
		res = append(res, float64(sum)/float64(sz))
		sum = 0
	}
	return res
}

// @lc code=end
